package grpc_gw

import (
	"context"
	"github.com/juicyluv/jiraya/internal/jiraya/interfaces/grpc_gw/adapters"
	"github.com/juicyluv/jiraya/internal/jiraya/interfaces/grpc_gw/protobuf"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) GetProject(ctx context.Context, request *protobuf.GetProjectRequest) (*protobuf.GetProjectResponse, error) {
	project, err := s.Core().GetProject(ctx, adapters.ToGetProjectRequest(request))

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return adapters.ToGetProjectResponse(project.Project), nil
}
