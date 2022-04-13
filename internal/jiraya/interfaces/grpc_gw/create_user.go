package grpc_gw

import (
	"context"
	"github.com/juicyluv/jiraya/internal/jiraya/domain"
	"github.com/juicyluv/jiraya/internal/jiraya/infrastructure/cast"
	"github.com/juicyluv/jiraya/internal/jiraya/interfaces/grpc_gw/protobuf"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) CreateUser(ctx context.Context, request *protobuf.CreateUserRequest) (*protobuf.GetUserResponse, error) {
	userID, err := s.Core().CreateUser(ctx, &domain.CreateUserRequest{
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	user, err := s.Core().GetUser(ctx, &domain.GetUserRequest{UserID: *userID})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &protobuf.GetUserResponse{User: &protobuf.User{
		Id:         user.UserID,
		Username:   user.Username,
		Email:      user.Email,
		CreatedAt:  cast.Timestamp(&user.CreatedAt),
		DisabledAt: cast.Timestamp(user.DisabledAt),
	}}, nil
}
