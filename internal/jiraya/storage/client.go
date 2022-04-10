package storage

import (
	"context"
	"github.com/juicyluv/jiraya/internal/jiraya/domain"
)

// Storage represents storage interface.
type Storage interface {
	CreateUser(ctx context.Context, request *domain.CreateUserRequest) (*domain.CreateUserResponse, error)
}
