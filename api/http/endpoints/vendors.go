package endpoints

import "github.com/labstack/echo/v4"

type Vendors struct {
}

func NewVendors() *Vendors {
	return &Vendors{}
}

func (h *Vendors) NewVendorsHandler(g *echo.Group) {
	vendorGroup := g.Group("/reports")

	vendorGroup.POST("", h.getDelayReport)
}

func (h *Vendors) getDelayReport(c echo.Context) error {
	return nil
}
