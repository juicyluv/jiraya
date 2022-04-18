package core

import (
	"context"
	"github.com/juicyluv/jiraya/internal/jiraya/domain"
)

func (c *Core) UpdateUserContact(ctx context.Context, request *domain.UpdateUserContactRequest) error {
	err := c.storage.UpdateUserContact(ctx, request)

	if err != nil {
		return err
	}

	return nil
}
