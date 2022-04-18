package grpc_gw

import (
	"context"
	"github.com/juicyluv/jiraya/internal/jiraya/interfaces/grpc_gw/adapters"
	"github.com/juicyluv/jiraya/internal/jiraya/interfaces/grpc_gw/protobuf"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) UpdateUserContact(ctx context.Context, request *protobuf.UpdateUserContactRequest) (*protobuf.UpdateUserContactResponse, error) {
	err := s.core.UpdateUserContact(ctx, adapters.ToUpdateUserContactRequest(request))

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &protobuf.UpdateUserContactResponse{}, nil
}
