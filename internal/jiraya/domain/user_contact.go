package domain

type UserContact struct {
	ContactID   string `json:"contact_id"`
	UserID      string `json:"user_id"`
	ContactName string `json:"contact_name"`
	Contact     string `json:"contact"`
}

type CreateUserContactRequest struct {
	UserID      string `json:"user_id"`
	ContactName string `json:"contact_name"`
	Contact     string `json:"contact"`
}

type GetUserContactRequest struct {
	ContactID string `json:"contact_id"`
}

type GetUserContactResponse struct {
	Contact *UserContact `json:"contact"`
}

type GetUserContactsRequest struct {
	UserID string `json:"user_id"`
}

type GetUserContactsResponse struct {
	Contacts []*UserContact `json:"contacts"`
}

type UpdateUserContactRequest struct {
	ContactID   string  `json:"contact_id"`
	ContactName *string `json:"contact_name"`
	Contact     *string `json:"contact"`
}

type DeleteUserContactRequest struct {
	ContactID string `json:"contact_id"`
}
