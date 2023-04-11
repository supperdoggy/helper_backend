package models

type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`

	FullName string `json:"full_name"`

	CreatedAt int64 `json:"created_at"`
	EditedAt  int64 `json:"edited_at"`
}
