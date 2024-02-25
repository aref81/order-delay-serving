package endpoints

import (
	"OrderDelayServing/pkg/repository"
	"github.com/labstack/echo/v4"
)

type Vendors struct {
	vendorRepo repository.VendorRepo
}

func NewVendors(vendorRepo repository.VendorRepo) *Vendors {
	return &Vendors{vendorRepo: vendorRepo}
}

func (h *Vendors) NewVendorsHandler(g *echo.Group) {
	vendorGroup := g.Group("/reports")

	vendorGroup.POST("", h.getDelayReport)
}

func (h *Vendors) getDelayReport(c echo.Context) error {
	return nil
}
