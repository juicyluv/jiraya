package postgres

import (
	"context"
	"github.com/juicyluv/jiraya/internal/jiraya/domain"
)

//language=PostgreSQL
const getUserContactSQL = `
	select
		id,
	    user_id,
	    contact_name,
	    contact
	from main.get_user_contact(
	    _contact_id := $1
	)
`

func (d *db) GetUserContact(ctx context.Context, request *domain.GetUserContactRequest) (*domain.GetUserContactResponse, error) {

	return nil, nil
}
