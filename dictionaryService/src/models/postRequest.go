package models

type PostRequest struct {
	Level  int    `json:"level,omitempty"`
	Text   string `json:"text,omitempty"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}
