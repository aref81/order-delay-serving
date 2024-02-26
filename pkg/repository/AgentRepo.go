package repository

import (
	"OrderDelayServing/pkg/model"
	"context"
	"gorm.io/gorm"
)

type AgentRepo interface {
	Create(ctx context.Context, agent model.Agent) (model.Agent, error)
	Get(ctx context.Context, agentID uint) (model.Agent, error)
	Update(ctx context.Context, agent model.Agent) error
	Delete(ctx context.Context, agentID uint) error
}

type AgentRepoImpl struct {
	db *gorm.DB
}

func NewAgentRepo(db *gorm.DB) *AgentRepoImpl {
	return &AgentRepoImpl{
		db: db,
	}
}

func (r *AgentRepoImpl) Create(ctx context.Context, agent model.Agent) (model.Agent, error) {
	result := r.db.WithContext(ctx).Create(&agent)
	if result.Error != nil {
		return model.Agent{}, result.Error
	}

	return agent, nil
}

func (r *AgentRepoImpl) Get(ctx context.Context, agentID uint) (model.Agent, error) {
	var agent model.Agent
	result := r.db.WithContext(ctx).Where(&model.Agent{ID: agentID}).First(&agent)
	if result.Error != nil {
		return model.Agent{}, result.Error
	}

	return agent, nil
}

func (r *AgentRepoImpl) Update(ctx context.Context, agent model.Agent) error {
	result := r.db.WithContext(ctx).Where(&model.Agent{ID: agent.ID}).Updates(&agent)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *AgentRepoImpl) Delete(ctx context.Context, agentID uint) error {
	result := r.db.WithContext(ctx).Where(&model.Agent{ID: agentID}).Delete(&model.Agent{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
