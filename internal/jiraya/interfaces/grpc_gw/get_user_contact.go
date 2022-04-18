package grpc_gw

import (
	"context"
	"github.com/juicyluv/jiraya/internal/jiraya/interfaces/grpc_gw/adapters"
	"github.com/juicyluv/jiraya/internal/jiraya/interfaces/grpc_gw/protobuf"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) GetUserContact(ctx context.Context, request *protobuf.GetUserContactRequest) (*protobuf.GetUserContactResponse, error) {
	contact, err := s.core.GetUserContact(ctx, adapters.ToGetUserContactRequest(request))

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return adapters.ToGetUserContactResponse(contact.Contact), nil
}
