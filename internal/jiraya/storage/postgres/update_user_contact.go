package postgres

import (
	"context"
	"github.com/juicyluv/jiraya/internal/jiraya/domain"
)

//language=PostgreSQL
const updateUserContactSQL = `
	select
	    update_user_contact
	from main.update_user_contact(
	    _contact_id := $1,
	    _contact_name := $2,
	    _contact := $3
	)
`

func (d *db) UpdateUserContact(ctx context.Context, request *domain.UpdateUserContactRequest) error {
	row := d.QueryRow(
		ctx,
		updateUserContactSQL,
		request.ContactID,
		request.ContactName,
		request.Contact,
	)

	var queryError []byte

	err := row.Scan(&queryError)

	if err != nil {
		return err
	}

	if err = handleQueryError(queryError); err != nil {
		return err
	}

	return nil
}
