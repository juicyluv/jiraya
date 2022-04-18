package core

import (
	"context"
	"github.com/juicyluv/jiraya/internal/jiraya/domain"
)

func (c *Core) GetUserContact(ctx context.Context, request *domain.GetUserContactRequest) (*domain.GetUserContactResponse, error) {
	contact, err := c.storage.GetUserContact(ctx, request)

	if err != nil {
		return nil, err
	}

	return &domain.GetUserContactResponse{Contact: contact}, nil
}
