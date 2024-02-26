package http

import (
	"OrderDelayServing/api/http/endpoints"
	"OrderDelayServing/internal/config"
	"OrderDelayServing/pkg/repository"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func Run(conf *config.Config, repos *repository.Repos) {
	e := echo.New()

	agents := endpoints.NewAgents(repos.AgentRepo)
	delayReports := endpoints.NewDelayReports(repos.DelayReportRepo)
	orders := endpoints.NewOrders(repos.OrderRepo)
	trips := endpoints.NewTrips(repos.TripRepo)
	vendors := endpoints.NewVendors(repos.VendorRepo, repos.OrderRepo)

	apiGroup := e.Group("/api")

	agents.NewAgentsHandler(apiGroup)
	delayReports.NewDelayReportsHandler(apiGroup)
	orders.NewOrdersHandler(apiGroup)
	trips.NewTripsHandler(apiGroup)
	vendors.NewVendorsHandler(apiGroup)

	if err := e.Start(conf.Server.Address + ":" + conf.Server.Port); err != nil {
		logrus.Fatalf("server failed to start %v", err)
	}
}
