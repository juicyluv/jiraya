package grpc_gw

import (
	"context"
	"github.com/juicyluv/jiraya/internal/jiraya/domain"
	"github.com/juicyluv/jiraya/internal/jiraya/infrastructure/cast"
	"github.com/juicyluv/jiraya/internal/jiraya/interfaces/grpc_gw/protobuf"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) GetUser(ctx context.Context, request *protobuf.GetUserRequest) (*protobuf.GetUserResponse, error) {
	user, err := s.Core().GetUser(ctx, &domain.GetUserRequest{UserID: request.UserId})

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
