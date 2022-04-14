package postgres

import (
	"context"
	"github.com/juicyluv/jiraya/internal/jiraya/domain"
)

//language=PostgreSQL
const createUserContactSQL = `
	select
		id,
	    error
	from main.create_user_contact(
	    _user_id := $1,
	    _contact_name := $2,
	    _contact := $3
	)
`

func (d *db) CreateUserContact(ctx context.Context, request *domain.CreateUserContactRequest) (*string, error) {
	row := d.QueryRow(
		ctx,
		createUserContactSQL,
		request.UserID,
		request.ContactName,
		request.Contact,
	)

	var (
		contactID  *string
		queryError []byte
	)

	err := row.Scan(&contactID, &queryError)

	if err != nil {
		return nil, err
	}

	if err = handleQueryError(queryError); err != nil {
		return nil, err
	}

	return contactID, nil
}
