package grpc_gw

import (
	"context"
	"github.com/juicyluv/jiraya/internal/jiraya/interfaces/grpc_gw/adapters"
	"github.com/juicyluv/jiraya/internal/jiraya/interfaces/grpc_gw/protobuf"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) DeleteUserContact(ctx context.Context, request *protobuf.DeleteUserContactRequest) (*protobuf.DeleteUserContactResponse, error) {
	err := s.core.DeleteUserContact(ctx, adapters.ToDeleteUserContactRequest(request))

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &protobuf.DeleteUserContactResponse{}, nil
}
