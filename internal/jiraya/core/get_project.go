package core

import (
	"context"
	"github.com/juicyluv/jiraya/internal/jiraya/domain"
)

func (c *Core) GetProject(ctx context.Context, request *domain.GetProjectRequest) (*domain.GetProjectResponse, error) {
	project, err := c.storage.GetProject(ctx, request)

	if err != nil {
		return nil, err
	}

	return &domain.GetProjectResponse{Project: project}, nil
}
