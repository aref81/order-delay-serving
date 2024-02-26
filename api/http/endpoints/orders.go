package endpoints

import (
	"OrderDelayServing/pkg/model"
	"OrderDelayServing/pkg/repository"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"
)

type Orders struct {
	orderRepo repository.OrderRepo
}

func NewOrders(orderRepo repository.OrderRepo) *Orders {
	return &Orders{orderRepo: orderRepo}
}

func (h *Orders) NewOrdersHandler(g *echo.Group) {
	ordersGroup := g.Group("/orders")

	ordersGroup.POST("", h.createNewOrder)
}

func (h *Orders) createNewOrder(c echo.Context) error {
	newOrder := new(model.Order)
	if err := c.Bind(newOrder); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	newOrder.RegisteredAt = time.Now()

	agent, err := h.orderRepo.Create(c.Request().Context(), *newOrder)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, agent)
}

func (h *Orders) getOrderByID(c echo.Context) error {
	oID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	orderID := uint(oID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	order, err := h.orderRepo.Get(c.Request().Context(), orderID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, order)
}
