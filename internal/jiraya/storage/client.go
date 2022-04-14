package storage

import (
	"context"
	"github.com/juicyluv/jiraya/internal/jiraya/domain"
)

// Storage represents storage interface.
type Storage interface {
	CreateUser(ctx context.Context, request *domain.CreateUserRequest) (*string, error)
	GetUser(ctx context.Context, request *domain.GetUserRequest) (*domain.User, error)
	GetUserByPassword(ctx context.Context, request *domain.GetUserByPasswordRequest) (*domain.User, error)
}
