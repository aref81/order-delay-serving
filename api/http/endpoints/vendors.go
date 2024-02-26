package endpoints

import (
	"OrderDelayServing/pkg/model"
	"OrderDelayServing/pkg/repository"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Vendors struct {
	vendorRepo repository.VendorRepo
	orderRepo  repository.OrderRepo
}

func NewVendors(vendorRepo repository.VendorRepo, orderRepo repository.OrderRepo) *Vendors {
	return &Vendors{
		vendorRepo: vendorRepo,
		orderRepo:  orderRepo,
	}
}

func (h *Vendors) NewVendorsHandler(g *echo.Group) {
	vendorsGroup := g.Group("/vendors")

	vendorsGroup.POST("", h.CreateNewVendor)
	vendorsGroup.GET("/:id", h.getVendorByID)
}

func (h *Vendors) CreateNewVendor(c echo.Context) error {
	newVendor := new(model.Vendor)
	if err := c.Bind(newVendor); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	agent, err := h.vendorRepo.Create(c.Request().Context(), *newVendor)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, agent)
}

func (h *Vendors) getVendorByID(c echo.Context) error {
	vID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	vendorID := uint(vID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	vendor, err := h.vendorRepo.Get(c.Request().Context(), vendorID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	vendor.Orders, err = h.orderRepo.GetOrdersByVendorID(c.Request().Context(), vendorID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, vendor)
}

func (h *Vendors) getDelayReport(c echo.Context) error {
	return nil
}
