package endpoints

import (
	"OrderDelayServing/pkg/model"
	"OrderDelayServing/pkg/repository"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Trips struct {
	tripRepo repository.TripRepo
}

func NewTrips(tripRepo repository.TripRepo) *Trips {
	return &Trips{tripRepo: tripRepo}
}

func (h *Trips) NewTripsHandler(g *echo.Group) {
	tripsGroup := g.Group("/trips")

	tripsGroup.POST("", h.createNewTrip)
	tripsGroup.GET("/:id", h.getTripByID)
}

func (h *Trips) createNewTrip(c echo.Context) error {
	newTrip := new(model.Trip)
	if err := c.Bind(newTrip); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	agent, err := h.tripRepo.Create(c.Request().Context(), *newTrip)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, agent)
}

func (h *Trips) getTripByID(c echo.Context) error {
	tID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	tripID := uint(tID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	trip, err := h.tripRepo.Get(c.Request().Context(), tripID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, trip)
}
