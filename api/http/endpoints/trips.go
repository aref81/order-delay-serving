package endpoints

import (
	"OrderDelayServing/pkg/repository"
	"github.com/labstack/echo/v4"
)

type Trips struct {
	tripRepo repository.TripRepo
}

func NewTrips(tripRepo repository.TripRepo) *Trips {
	return &Trips{tripRepo: tripRepo}
}

func (h *Trips) NewTripsHandler(g *echo.Group) {
	return
}
