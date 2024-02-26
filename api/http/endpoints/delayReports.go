package endpoints

import (
	"OrderDelayServing/pkg/model"
	"OrderDelayServing/pkg/repository"
	"OrderDelayServing/utils/broker"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

type DelayResponse struct {
	Delay int `json:"delay"`
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
	reportsGroup.GET("", h.getQueuedReport)
}

func (h *DelayReports) reportDelay(c echo.Context) error {
	newReport := new(model.DelayReport)
	if err := c.Bind(newReport); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	trip, DBerr := h.tripRepo.GetByOrderID(c.Request().Context(), newReport.OrderID)
	if DBerr != nil {
		if !errors.Is(DBerr, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": DBerr.Error()})
		}
	} else if trip.Status == model.TripStatusAssigned || trip.Status == model.TripStatusPicked ||
		trip.Status == model.TripStatusAtVendor {
		newDelay, err := requestNewDelay()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to call external API"})
		}

		newReport.DelayAmount = newDelay
	}

	report, err := h.delayReportRepo.Create(c.Request().Context(), *newReport)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": DBerr.Error()})
	}

	if errors.Is(DBerr, gorm.ErrRecordNotFound) || !(trip.Status == model.TripStatusAssigned || trip.Status == model.TripStatusPicked ||
		trip.Status == model.TripStatusAtVendor) {
		err := h.enqueueReport(&report)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to send report ID to RabbitMQ"})
		}
	}

	return c.JSON(http.StatusCreated, report)
}

func (h *DelayReports) getQueuedReport(c echo.Context) error {
	return nil
}

func (h *DelayReports) enqueueReport(report *model.DelayReport) error {
	body := strconv.FormatUint(uint64(report.ID), 10)
	err := h.rabbit.Channel.Publish(
		"",                  // exchange
		h.rabbit.Queue.Name, // routing key (queue name)
		false,               // mandatory
		false,               // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		return err
	}

	return nil
}

func requestNewDelay() (time.Duration, error) {
	resp, err := http.Get("https://run.mocky.io/v3/9f1ba24e-a43f-448a-bcde-c0737a84093f")
	if err != nil {
		return 0, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("API request returned status code: %d", resp.StatusCode)
	}

	var delayResp DelayResponse
	if err := json.NewDecoder(resp.Body).Decode(&delayResp); err != nil {
		return 0, fmt.Errorf("error decoding response: %w", err)
	}

	delayDuration := time.Duration(delayResp.Delay) * time.Second

	return delayDuration, nil
}
