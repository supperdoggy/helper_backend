package models

// Advert is a model for advert
type Advert struct {
	UserID      string   `json:"user_id"`
	Name        string   `json:"name"`
	Body        string   `json:"body"`
	Type        string   `json:"type"`
	Category    string   `json:"category"`
	Location    string   `json:"location"`
	Attachments []string `json:"attachments"`

	CreatedAt int64 `json:"created_at"`
	EditedAt  int64 `json:"edited_at"`
}

type Adverts struct {
	Adverts []Advert `json:"adverts"`
}

type AdvertsFilter struct {
	UserID    *string `json:"user_id,omitempty"`
	Type      *string `json:"type,omitempty"`
	Category  *string `json:"category,omitempty"`
	Location  *string `json:"location,omitempty"`
	Name      *string `json:"name,omitempty"`
	CreatedAt *int64  `json:"created_at,omitempty"`
}
