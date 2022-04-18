package postgres

import (
	"context"
	"github.com/juicyluv/jiraya/internal/jiraya/domain"
)

//language=PostgreSQL
const createProjectSQL = `
	select 
		project_id,
	    error
	from main.create_project(
	    _title := $1,
	    _creator_id := $2,
	    _description := $3,
	    _icon_url := $4
	)
`

func (d *db) CreateProject(ctx context.Context, request *domain.CreateProjectRequest) (*string, error) {
	row := d.QueryRow(
		ctx,
		createProjectSQL,
		request.Title,
		request.CreatorID,
		request.Description,
		request.IconURL,
	)

	var (
		projectID  string
		queryError []byte
	)

	err := row.Scan(&projectID, &queryError)

	if err != nil {
		return nil, err
	}

	if err = handleQueryError(queryError); err != nil {
		return nil, err
	}

	return &projectID, nil
}
