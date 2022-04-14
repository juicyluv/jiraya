package postgres

import (
	"context"
	"github.com/juicyluv/jiraya/internal/jiraya/domain"
)

//language=PostgreSQL
const getUserContactsSQL = `
	select
	    get_user_contacts
	from main.get_user_contacts(
	    _user_id := $1
	)
`

func (d *db) GetUserContacts(ctx context.Context, request *domain.GetUserContactsRequest) ([]*domain.UserContact, error) {
	rows, err := d.Query(ctx, getUserContactsSQL, request.UserID)

	if err != nil {
		return nil, err
	}

	contacts := make([]*domain.UserContact, 0)

	for rows.Next() {
		contact := &domain.UserContact{}

		err := rows.Scan(
			&contact.ContactID,
			&contact.ContactName,
			&contact.Contact,
		)

		if err != nil {
			return nil, err
		}

		contact.UserID = request.UserID
		contacts = append(contacts, contact)
	}

	return contacts, nil
}
