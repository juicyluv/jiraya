package postgres

import (
	"context"
	"github.com/juicyluv/jiraya/internal/jiraya/domain"
)

//language=PostgreSQL
const getProjectSQL = `
	select
	    title,
	    creator_id,
	    description,
	    icon_url,
	    created_at,
	    closed_at,
	    error
	from main.get_project(
	    _project_id := $1
	)
`

func (d *db) GetProject(ctx context.Context, request *domain.GetProjectRequest) (*domain.Project, error) {
	row := d.QueryRow(ctx, getProjectSQL, request.ProjectID)

	var (
		project    domain.Project
		queryError []byte
	)

	err := row.Scan(
		&project.Title,
		&project.CreatorID,
		&project.Description,
		&project.IconURL,
		&project.CreatedAt,
		&project.ClosedAt,
		&queryError,
	)

	if err != nil {
		return nil, err
	}

	if err = handleQueryError(queryError); err != nil {
		return nil, err
	}

	project.ProjectID = request.ProjectID

	return &project, nil
}
