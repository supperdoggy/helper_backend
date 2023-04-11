package models

type CreateUserRequest struct {
	Password string `json:"password"`
	Email    string `json:"email"`
	FullName string `json:"fill_name"`
}

type CreateUserResponse struct {
	ID    string `json:"id,omitempty"`
	Token string `json:"token,omitempty"`
	Error string `json:"error,omitempty"`
}

type DeleteUserRequest struct {
	ID string `json:"id,omitempty"`
}
type DeleteUserResponse struct {
	ID    *string `json:"id,omitempty"`
	Error string  `json:"error,omitempty"`
}

type UpdateUserRequest struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserResponse struct {
	User  *User  `json:"user,omitempty"`
	Error string `json:"error,omitempty"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type GetUserResponse struct {
	User  *User  `json:"user,omitempty"`
	Error string `json:"error,omitempty"`
}
