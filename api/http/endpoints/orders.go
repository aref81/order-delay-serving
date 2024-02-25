package endpoints

import (
	"OrderDelayServing/pkg/repository"
	"github.com/labstack/echo/v4"
)

type Orders struct {
	orderRepo repository.OrderRepo
}

func NewOrders(orderRepo repository.OrderRepo) *Orders {
	return &Orders{orderRepo: orderRepo}
}

func (h *Orders) NewOrdersHandler(g *echo.Group) {
	return
}
