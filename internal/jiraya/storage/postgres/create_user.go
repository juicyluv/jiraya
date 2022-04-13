package postgres

import (
	"context"
	"errors"
	"github.com/juicyluv/jiraya/internal/jiraya/domain"
)

//language=PostgreSQL
const createUserSQL = `
	select 
		id,
	    error
	from main.create_user(
		_username := $1,
	    _email := $2,
	    _password := $3
	)
`

func (d *db) CreateUser(ctx context.Context, request *domain.CreateUserRequest) (*string, error) {
	rows, err := d.Query(ctx, createUserSQL,
		request.Username,
		request.Email,
		request.Password,
	)

	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, errors.New("null response")
	}

	var (
		userID     *string
		queryError []byte
	)

	err = rows.Scan(&userID, &queryError)

	if err != nil {
		return nil, err
	}

	if err = handleQueryError(queryError); err != nil {
		return nil, err
	}

	return userID, nil
}
