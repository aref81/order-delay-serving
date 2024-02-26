package endpoints

import (
	"OrderDelayServing/pkg/model"
	"OrderDelayServing/pkg/repository"
	"OrderDelayServing/utils/broker"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

type DelayResponse struct {
	Delay int `json:"delay"`
}

type ReviewSubmission struct {
	ReportID uint          `json:"reportID"`
	AgentID  uint          `json:"agentID"`
	Status   string        `json:"status"`
	Delay    time.Duration `json:"deliveryTime"`
}

type DelayReports struct {
	agentRepo       repository.AgentRepo
	delayReportRepo repository.DelayReportRepo
	orderRepo       repository.OrderRepo
	tripRepo        repository.TripRepo
	vendorRepo      repository.VendorRepo
	rabbit          broker.RabbitMQ
}

func NewDelayReports(agentRepo repository.AgentRepo, delayReportRepo repository.DelayReportRepo,
	orderRepo repository.OrderRepo, tripRepo repository.TripRepo, vendorRepo repository.VendorRepo,
	rabbit broker.RabbitMQ) *DelayReports {
	return &DelayReports{
		agentRepo:       agentRepo,
		delayReportRepo: delayReportRepo,
		orderRepo:       orderRepo,
		tripRepo:        tripRepo,
		vendorRepo:      vendorRepo,
		rabbit:          rabbit,
	}
}

func (h *DelayReports) NewDelayReportsHandler(g *echo.Group) {
	reportsGroup := g.Group("/reports")

	reportsGroup.POST("", h.reportDelay)
	reportsGroup.POST("/:agentID", h.getQueuedReport)
	reportsGroup.GET("", h.getVendorsSummary)
	reportsGroup.POST("/review", h.submitReview)
}

func (h *DelayReports) reportDelay(c echo.Context) error {
	newReport := new(model.DelayReport)
	var msg string

	if err := c.Bind(newReport); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	order, err := h.orderRepo.Get(c.Request().Context(), newReport.OrderID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	trip, DBerr := h.tripRepo.GetByOrderID(c.Request().Context(), newReport.OrderID)
	if DBerr != nil {
		if !errors.Is(DBerr, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": DBerr.Error()})
		}
	}

	if order.RegisteredAt.Add(order.DeliveryTime).After(time.Now()) {
		return c.JSON(http.StatusOK, map[string]string{"msg": "can't report delay before a trip's delivery time"})
	}

	if trip.Status == model.TripStatusQueued || trip.Status == model.TripStatusOnReview {
		return c.JSON(http.StatusOK, map[string]string{"msg": "this trip is still under review"})
	}

	if trip.Status == model.TripStatusAssigned || trip.Status == model.TripStatusPicked ||
		trip.Status == model.TripStatusAtVendor {
		newDelay, err := requestNewDelay()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to call external API"})
		}

		newReport.DelayAmount = newDelay
		order.DeliveryTime = order.DeliveryTime + newDelay
		msg = fmt.Sprintf("The order will be delivered to you in %f minutes", newDelay.Minutes())
	}

	newReport.IssuedAt = time.Now()
	report, err := h.delayReportRepo.Create(c.Request().Context(), *newReport)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": DBerr.Error()})
	}

	if errors.Is(DBerr, gorm.ErrRecordNotFound) || !(trip.Status == model.TripStatusAssigned ||
		trip.Status == model.TripStatusPicked || trip.Status == model.TripStatusAtVendor) {
		err := h.enqueueReport(&report)
		trip.Status = model.TripStatusQueued
		msg = fmt.Sprintf("this order is queued for review")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to send report ReportID to RabbitMQ"})
		}
	}

	if !errors.Is(DBerr, gorm.ErrRecordNotFound) {
		err = h.tripRepo.Update(c.Request().Context(), trip)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
	}

	err = h.orderRepo.Update(c.Request().Context(), order)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"msg": msg})
}

func (h *DelayReports) getQueuedReport(c echo.Context) error {
	aID, err := strconv.ParseUint(c.Param("agentID"), 10, 32)
	agentID := uint(aID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	agent, err := h.agentRepo.Get(c.Request().Context(), agentID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	if !agent.IsAvailable {
		return c.JSON(http.StatusOK, map[string]string{"msg": "can't assign more than one report to any agent"})
	}

	reportID, err := h.dequeueReport()
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	report, err := h.delayReportRepo.Get(c.Request().Context(), reportID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	trip, tripErr := h.tripRepo.GetByOrderID(c.Request().Context(), report.OrderID)
	if tripErr != nil {
		logrus.Warnf("error: %v", err)
	}

	if tripErr == nil {
		if trip.Status != model.TripStatusQueued {
			return c.JSON(http.StatusInternalServerError, map[string]string{"msg": "error while fetching report, please try again"})
		}

		trip.Status = model.TripStatusOnReview

		err = h.tripRepo.Update(c.Request().Context(), trip)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
	}

	agent.CurrentReportID = reportID
	agent.IsAvailable = false
	err = h.agentRepo.Update(c.Request().Context(), agent)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, report)
}

func (h *DelayReports) getVendorsSummary(c echo.Context) error {
	vendorsSummary, err := h.delayReportRepo.GetVendorsSummary(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, vendorsSummary)
}

func (h *DelayReports) submitReview(c echo.Context) error {
	submission := new(ReviewSubmission)
	if err := c.Bind(submission); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	report, err := h.delayReportRepo.Get(c.Request().Context(), submission.ReportID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	trip, tripErr := h.tripRepo.GetByOrderID(c.Request().Context(), report.OrderID)
	if tripErr != nil {
		logrus.Warnf("error: %v", err)
	}

	if tripErr == nil {
		if trip.Status != model.TripStatusOnReview {
			return c.JSON(http.StatusInternalServerError, map[string]string{"msg": "this report is not under review"})
		}
	}

	order, err := h.orderRepo.Get(c.Request().Context(), trip.OrderID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	trip.Status = submission.Status
	order.DeliveryTime = order.DeliveryTime + submission.Delay
	report.DelayAmount = submission.Delay

	if tripErr == nil {
		err = h.tripRepo.Update(c.Request().Context(), trip)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
	}

	err = h.orderRepo.Update(c.Request().Context(), order)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	err = h.delayReportRepo.Update(c.Request().Context(), report)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"msg": "review submitted successfully"})
}

func (h *DelayReports) enqueueReport(report *model.DelayReport) error {
	body := strconv.FormatUint(uint64(report.ID), 10)
	err := h.rabbit.Channel.Publish(
		"",
		h.rabbit.Queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		return err
	}

	return nil
}

func (h *DelayReports) dequeueReport() (uint, error) {
	msg, ok, err := h.rabbit.Channel.Get(
		h.rabbit.Queue.Name,
		true,
	)
	if err != nil {
		return 0, err
	}

	if !ok {
		return 0, errors.New("empty queue")
	}

	rID, err := strconv.ParseUint(string(msg.Body), 10, 32)

	return uint(rID), nil
}

func requestNewDelay() (time.Duration, error) {
	resp, err := http.Get("https://run.mocky.io/v3/9f1ba24e-a43f-448a-bcde-c0737a84093f")
	if err != nil {
		return 0, errors.New("error making request: " + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, errors.New("API request returned status: " + resp.Status)
	}

	var delayResp DelayResponse
	if err := json.NewDecoder(resp.Body).Decode(&delayResp); err != nil {
		return 0, errors.New("error decoding response: " + err.Error())
	}

	delayDuration := time.Duration(delayResp.Delay) * time.Second

	return delayDuration, nil
}
