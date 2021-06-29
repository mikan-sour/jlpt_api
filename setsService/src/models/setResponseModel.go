package models

type SetResponse struct {
	StatusCode int   `json:"statusCode"`
	Sets       []Set `json:"sets"`
}
