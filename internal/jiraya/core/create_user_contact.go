package core

import (
	"context"
	"github.com/juicyluv/jiraya/internal/jiraya/domain"
)

func (c *Core) CreateUserContact(ctx context.Context, request *domain.CreateUserContactRequest) (*domain.GetUserContactResponse, error) {
	contactID, err := c.storage.CreateUserContact(ctx, request)

	if err != nil {
		return nil, err
	}

	contact, err := c.storage.GetUserContact(ctx, &domain.GetUserContactRequest{ContactID: *contactID})

	if err != nil {
		return nil, err
	}

	return &domain.GetUserContactResponse{Contact: contact}, nil
}
