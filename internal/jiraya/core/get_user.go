package core

import (
	"context"
	"github.com/juicyluv/jiraya/internal/jiraya/domain"
)

func (c *Core) GetUser(ctx context.Context, request *domain.GetUserRequest) (*domain.User, error) {
	user, err := c.storage.GetUser(ctx, request)

	return user, err
}
