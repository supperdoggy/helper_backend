package dbmodels

import "time"

type User struct {
	ID       string `bson:"_id"`
	Email    string `bson:"email"`
	FullName string `bson:"full_name"`
	Password []byte `bson:"password"`

	CreatedAt int64 `bson:"created_at"`
	EditedAt  int64 `bson:"edited_at"`
}

type Token struct {
	UserID   string    `json:"user_id"`
	TokenStr string    `json:"token_str"`
	Expire   time.Time `json:"expire"`
}
