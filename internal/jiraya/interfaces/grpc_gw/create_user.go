package grpc_gw

import (
	"context"
	"github.com/juicyluv/jiraya/internal/jiraya/interfaces/grpc_gw/protobuf"
)

func (s *server) CreateUser(ctx context.Context, request *protobuf.CreateUserRequest) (*protobuf.CreateUserResponse, error) {
	return &protobuf.CreateUserResponse{Message: "hello, world!"}, nil
}
