package postgres

import (
	"context"
	"github.com/juicyluv/jiraya/internal/jiraya/domain"
)

//language=PostgreSQL
const getUserContactSQL = `
	select
	    user_id,
	    contact_name,
	    contact,
	    error
	from main.get_user_contact(
	    _contact_id := $1
	)
`

func (d *db) GetUserContact(ctx context.Context, request *domain.GetUserContactRequest) (*domain.UserContact, error) {
	row := d.QueryRow(ctx, getUserContactSQL, request.ContactID)

	var (
		contact    domain.UserContact
		queryError []byte
	)

	err := row.Scan(
		&contact.UserID,
		&contact.ContactName,
		&contact.Contact,
		&queryError,
	)

	if err != nil {
		return nil, err
	}

	if err = handleQueryError(queryError); err != nil {
		return nil, err
	}

	contact.ContactID = request.ContactID

	return &contact, nil
}
