package models

type CreateAdvertRequest struct {
	UserID      string   `json:"user_id"`
	Name        string   `json:"name"`
	Body        string   `json:"body"`
	Type        string   `json:"type"`
	Category    string   `json:"category"`
	Location    string   `json:"location"`
	Attachments [][]byte `json:"attachments"`
}

type CreateAdvertResponse struct {
	Advert *Advert `json:"advert"`
	Error  string  `json:"error"`
}

type DeleteAdvertRequest struct {
	ID string `json:"id"`
}

type DeleteAdvertResponse struct {
	ID    string `json:"id"`
	Error string `json:"error"`
}

type GetAdvertRequest struct {
	ID string `json:"id"`
}

type GetAdvertResponse struct {
	Advert *Advert `json:"advert"`
	Error  string  `json:"error"`
}

type GetAdvertsRequest struct {
	Filter *AdvertsFilter `json:"filter"`
	Limit  int            `json:"limit"`
	Offset int            `json:"offset"`
}

type GetAdvertsResponse struct {
	Adverts []Advert `json:"adverts"`
	Error   string   `json:"error"`
}
