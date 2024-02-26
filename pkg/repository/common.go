package repository

import "gorm.io/gorm"

type Repos struct {
	AgentRepo       *AgentRepoImpl
	DelayReportRepo *DelayReportRepoImpl
	OrderRepo       *OrderRepoImpl
	TripRepo        *TripRepoImpl
	VendorRepo      *VendorRepoImpl
}

func InitRepos(db *gorm.DB) *Repos {
	return &Repos{
		AgentRepo:       NewAgentRepo(db),
		DelayReportRepo: NewDelayReportRepo(db),
		OrderRepo:       NewOrderRepo(db),
		TripRepo:        NewTripRepo(db),
		VendorRepo:      NewVendorRepo(db),
	}
}
