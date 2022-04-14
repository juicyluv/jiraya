package postgres

import (
	"context"
	"github.com/juicyluv/jiraya/internal/jiraya/domain"
)

//language=PostgreSQL
const getUserByPasswordSQL = `
	select
		id,
	    username,
	    email,
	    created_at,
	    disabled_at,
	    error
	from main.get_user_by_password(
	    _login := $1,
	    _password := $2
	)
`

func (d *db) GetUserByPassword(ctx context.Context, request *domain.GetUserByPasswordRequest) (*domain.User, error) {
	row := d.QueryRow(ctx, getUserByPasswordSQL, request.Login, request.Password)

	var (
		user       *domain.User
		queryError []byte
	)

	err := row.Scan(
		&user.UserID,
		&user.Username,
		&user.Email,
		&user.CreatedAt,
		&user.DisabledAt,
		queryError,
	)

	if err != nil {
		return nil, err
	}

	if err = handleQueryError(queryError); err != nil {
		return nil, err
	}

	return user, nil
}
