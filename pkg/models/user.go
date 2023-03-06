package models

type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`

	CreatedAt int64 `json:"created_at"`
	EditedAt  int64 `json:"edited_at"`
}
