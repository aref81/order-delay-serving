package endpoints

import (
	"OrderDelayServing/pkg/repository"
	"github.com/labstack/echo/v4"
)

type Agents struct {
	agentRepo repository.AgentRepo
}

func NewAgents(agentRepo repository.AgentRepo) *Agents {
	return &Agents{agentRepo: agentRepo}
}

func (h *Agents) NewAgentsHandler(g *echo.Group) {
	return
}
