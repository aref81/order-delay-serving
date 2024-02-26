package endpoints

import (
	"OrderDelayServing/pkg/model"
	"OrderDelayServing/pkg/repository"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Agents struct {
	agentRepo repository.AgentRepo
}

func NewAgents(agentRepo repository.AgentRepo) *Agents {
	return &Agents{agentRepo: agentRepo}
}

func (h *Agents) NewAgentsHandler(g *echo.Group) {
	agentsGroup := g.Group("/agents")

	agentsGroup.POST("", h.createNewAgent)
}

func (h *Agents) createNewAgent(c echo.Context) error {
	newAgent := new(model.Agent)
	if err := c.Bind(newAgent); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	agent, err := h.agentRepo.Create(c.Request().Context(), *newAgent)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, agent)
}
