package core

import (
	"context"
	"github.com/juicyluv/jiraya/internal/jiraya/domain"
)

func (c *Core) CreateProject(ctx context.Context, request *domain.CreateProjectRequest) (*domain.GetProjectResponse, error) {
	projectID, err := c.storage.CreateProject(ctx, request)

	if err != nil {
		return nil, err
	}

	project, err := c.storage.GetProject(ctx, &domain.GetProjectRequest{ProjectID: *projectID})

	if err != nil {
		return nil, err
	}

	return &domain.GetProjectResponse{Project: project}, nil
}
