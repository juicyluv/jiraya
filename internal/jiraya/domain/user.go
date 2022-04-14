package domain

import "time"

type User struct {
	UserID     string     `json:"user_id"`
	Username   string     `json:"username"`
	Email      string     `json:"email"`
	CreatedAt  time.Time  `json:"created_at"`
	DisabledAt *time.Time `json:"disabled_at"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetUserRequest struct {
	UserID string `json:"user_id"`
}

type GetUserResponse struct {
	User *User `json:"user"`
}

type GetUserByPasswordRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
