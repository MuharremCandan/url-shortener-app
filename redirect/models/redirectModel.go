package models

type Redirect struct {
	Code      string `json:"code" `
	URL       string `json:"url"`
	CreatedAt int64  `json:"created_at"`
}
