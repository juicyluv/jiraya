package postgres

import (
	"context"
	"github.com/juicyluv/jiraya/internal/jiraya/domain"
)

//language=PostgreSQL
const getUserSQL = `
	select 
	    username,
	    email,
	    created_at,
	    disabled_at,
	    error
	from main.get_user(
	    _id := $1
	)
`

func (d *db) GetUser(ctx context.Context, request *domain.GetUserRequest) (*domain.User, error) {
	rows, err := d.Query(ctx, getUserSQL, request.UserID)

	if err != nil {
		return nil, err
	}

	var (
		user       domain.User
		queryError []byte
	)

	err = rows.Scan(
		&user.Username,
		&user.Email,
		&user.CreatedAt,
		&user.DisabledAt,
		&queryError,
	)

	if err != nil {
		return nil, err
	}

	if err = handleQueryError(queryError); err != nil {
		return nil, err
	}

	return &user, nil
}
