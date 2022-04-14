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
