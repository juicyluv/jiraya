package core

import (
	"context"
	"github.com/juicyluv/jiraya/internal/jiraya/domain"
)

func (c *Core) DeleteUserContact(ctx context.Context, request *domain.DeleteUserContactRequest) error {
	err := c.storage.DeleteUserContact(ctx, request)

	if err != nil {
		return err
	}

	return nil
}
