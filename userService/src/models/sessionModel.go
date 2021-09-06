package models

import (
	"database/sql"
	"fmt"
	"time"
)

type Session struct {
	ID          int       `json:"sessionId"`
	OwnerId     int       `json:"ownerId"`
	CreatedDate time.Time `json:"createdDate"`
}

func (s *Session) CreateSession(DB *sql.DB) (*Session, *ApiError) {

	var insertSessionQuery = "INSERT INTO sessions(owner_id) VALUES(%v) RETURNING owner_id;"

	err := DB.QueryRow(fmt.Sprintf(insertSessionQuery, s.OwnerId)).Scan(&s.OwnerId)

	if err != nil {
		return nil, &ApiError{ErrorMessage: err.Error()}
	}

	return &Session{s.ID, s.OwnerId, s.CreatedDate}, nil
}
