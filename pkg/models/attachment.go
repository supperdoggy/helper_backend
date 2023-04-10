package models

// Attachment is a model for attachment
type Attachment struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Data []byte `json:"data"`
}
