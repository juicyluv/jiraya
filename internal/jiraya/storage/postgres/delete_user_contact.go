package postgres

import (
	"context"
	"github.com/juicyluv/jiraya/internal/jiraya/domain"
)

//language=PostgreSQL
const deleteUserContactSQL = `
	select
	    delete_user_contact
	from main.delete_user_contact(
	    _contact_id := $1
	)
`

func (d *db) DeleteUserContact(ctx context.Context, request *domain.DeleteUserContactRequest) error {
	row := d.QueryRow(ctx, deleteUserContactSQL, request.ContactID)

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
