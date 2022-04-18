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

	CreateUserContact(ctx context.Context, request *domain.CreateUserContactRequest) (*string, error)
	GetUserContact(ctx context.Context, request *domain.GetUserContactRequest) (*domain.UserContact, error)
	GetUserContacts(ctx context.Context, request *domain.GetUserContactsRequest) ([]*domain.UserContact, error)
	UpdateUserContact(ctx context.Context, request *domain.UpdateUserContactRequest) error
	DeleteUserContact(ctx context.Context, request *domain.DeleteUserContactRequest) error

	CreateProject(ctx context.Context, request *domain.CreateProjectRequest) (*string, error)
	GetProject(ctx context.Context, request *domain.GetProjectRequest) (*domain.Project, error)
}
