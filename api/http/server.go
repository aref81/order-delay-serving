package http

import (
	"OrderDelayServing/api/http/endpoints"
	"OrderDelayServing/internal/config"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func Run(conf *config.Config) {
	e := echo.New()

	dr := endpoints.NewDelayReports()
	v := endpoints.NewVendors()

	apiGroup := e.Group("/api")

	dr.NewDelayReportsHandler(apiGroup)
	v.NewVendorsHandler(apiGroup)

	if err := e.Start(conf.Server.Address + ":" + conf.Server.Port); err != nil {
		logrus.Fatalf("server failed to start %v", err)
	}
}
