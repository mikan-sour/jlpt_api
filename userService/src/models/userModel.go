package models

import "time"

// this is used to get the user from the DB

type User struct {
	ID           int       `json:"id"`
	Username     string    `json:"username"`
	IsAdmin      bool      `json:"isAdmin"`
	Active       bool      `json:"active"`
	CreatedDate  time.Time `json:"createdDate"`
	LastLogin    time.Time `json:"lastLogin"`
	ModifiedDate time.Time `json:"modifiedDate"`
}
