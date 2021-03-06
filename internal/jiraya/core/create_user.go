package core

import (
	"context"
	"github.com/juicyluv/jiraya/internal/jiraya/domain"
)

func (c *Core) CreateUser(ctx context.Context, request *domain.CreateUserRequest) (*string, error) {
	userID, err := c.storage.CreateUser(ctx, request)

	return userID, err
}
