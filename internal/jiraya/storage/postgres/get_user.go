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
	    _user_id := $1
	)
`

func (d *db) GetUser(ctx context.Context, request *domain.GetUserRequest) (*domain.User, error) {
	row := d.QueryRow(ctx, getUserSQL, request.UserID)

	var (
		user       *domain.User
		queryError []byte
	)

	err := row.Scan(
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

	user.UserID = request.UserID

	return user, nil
}
