package endpoints

import (
	"OrderDelayServing/pkg/repository"
	"github.com/labstack/echo/v4"
)

type DelayReports struct {
	delayReportRepo repository.DelayReportRepo
}

func NewDelayReports(DelayReportRepo repository.DelayReportRepo) *DelayReports {
	return &DelayReports{delayReportRepo: DelayReportRepo}
}

func (h *DelayReports) NewDelayReportsHandler(g *echo.Group) {
	reportsGroup := g.Group("/reports")

	reportsGroup.POST("", h.reportDelay)
	reportsGroup.GET("", h.getQueuedReport)
}

func (h *DelayReports) reportDelay(c echo.Context) error {
	return nil
}

func (h *DelayReports) getQueuedReport(c echo.Context) error {
	return nil
}
