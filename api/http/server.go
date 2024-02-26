package http

import (
	"OrderDelayServing/api/http/endpoints"
	"OrderDelayServing/internal/config"
	"OrderDelayServing/pkg/repository"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type repos struct {
	agentRepo       *repository.AgentRepoImpl
	delayReportRepo *repository.DelayReportRepoImpl
	orderRepo       *repository.OrderRepoImpl
	tripRepo        *repository.TripRepoImpl
	vendorRepo      *repository.VendorRepoImpl
}

func initRepos(db *gorm.DB) *repos {
	return &repos{
		agentRepo:       repository.NewAgentRepo(db),
		delayReportRepo: repository.NewDelayReportRepo(db),
		orderRepo:       repository.NewOrderRepo(db),
		tripRepo:        repository.NewTripRepo(db),
		vendorRepo:      repository.NewVendorRepo(db),
	}
}

func Run(conf *config.Config, db *gorm.DB) {
	repos := initRepos(db)

	e := echo.New()

	agents := endpoints.NewAgents(repos.agentRepo)
	delayReports := endpoints.NewDelayReports(repos.delayReportRepo)
	orders := endpoints.NewOrders(repos.orderRepo)
	trips := endpoints.NewTrips(repos.tripRepo)
	vendors := endpoints.NewVendors(repos.vendorRepo, repos.orderRepo)

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
