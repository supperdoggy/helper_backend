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

type EmailCode struct {
	Email  string    `json:"email"`
	Code   string    `json:"code"`
	Expire time.Time `json:"expire"`
}

type Advert struct {
	ID          string   `bson:"_id"`
	UserID      string   `bson:"user_id"`
	Name        string   `bson:"name"`
	Body        string   `bson:"body"`
	Type        string   `bson:"type"`
	Category    string   `bson:"category"`
	Location    string   `bson:"location"`
	Attachments []string `bson:"attachments"`

	CreatedAt int64 `bson:"created_at"`
	EditedAt  int64 `bson:"edited_at"`
}

type Attachment struct {
	ID   string `bson:"_id"`
	Name string `bson:"name"`
	Data []byte `bson:"data"`

	CreatedAt int64 `bson:"created_at"`
	EditedAt  int64 `bson:"edited_at"`
}
