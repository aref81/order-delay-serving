package endpoints

import "github.com/labstack/echo/v4"

type DelayReports struct {
}

func NewDelayReports() *DelayReports {
	return &DelayReports{}
}

func (h *DelayReports) NewDelayReportsHandler(g *echo.Group) {
	reportGroup := g.Group("/reports")

	reportGroup.POST("", h.reportDelay)
	reportGroup.GET("", h.getQueuedReport)
}

func (h *DelayReports) reportDelay(c echo.Context) error {
	return nil
}

func (h *DelayReports) getQueuedReport(c echo.Context) error {
	return nil
}
