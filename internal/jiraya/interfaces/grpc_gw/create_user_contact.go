package grpc_gw

import (
	"context"
	"github.com/juicyluv/jiraya/internal/jiraya/interfaces/grpc_gw/adapters"
	"github.com/juicyluv/jiraya/internal/jiraya/interfaces/grpc_gw/protobuf"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) CreateUserContact(ctx context.Context, request *protobuf.CreateUserContactRequest) (*protobuf.GetUserContactResponse, error) {
	contact, err := s.core.CreateUserContact(ctx, adapters.ToCreateUserContactRequest(request))

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return adapters.ToGetUserContactResponse(contact.Contact), nil
}
