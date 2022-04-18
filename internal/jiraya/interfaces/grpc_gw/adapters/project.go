package adapters

import (
	"github.com/juicyluv/jiraya/internal/jiraya/domain"
	"github.com/juicyluv/jiraya/internal/jiraya/infrastructure/cast"
	"github.com/juicyluv/jiraya/internal/jiraya/interfaces/grpc_gw/protobuf"
)

func ToProject(project *domain.Project) *protobuf.Project {
	if project == nil {
		return nil
	}

	return &protobuf.Project{
		ProjectId:   project.ProjectID,
		Title:       project.Title,
		CreatorId:   project.CreatorID,
		Description: project.Description,
		IconUrl:     project.IconURL,
		CreatedAt:   cast.Timestamp(&project.CreatedAt),
		ClosedAt:    cast.Timestamp(project.ClosedAt),
	}
}

func ToGetProjectRequest(req *protobuf.GetProjectRequest) *domain.GetProjectRequest {
	if req == nil {
		return nil
	}

	return &domain.GetProjectRequest{ProjectID: req.ProjectId}
}

func ToCreateProjectRequest(req *protobuf.CreateProjectRequest) *domain.CreateProjectRequest {
	if req == nil {
		return nil
	}

	return &domain.CreateProjectRequest{
		Title:       req.Title,
		CreatorID:   req.CreatorId,
		Description: req.Description,
		IconURL:     req.IconUrl,
	}
}

func ToGetProjectResponse(project *domain.Project) *protobuf.GetProjectResponse {
	if project == nil {
		return nil
	}

	return &protobuf.GetProjectResponse{
		Project: ToProject(project),
	}
}
