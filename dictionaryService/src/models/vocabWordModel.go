package models

import (
	"database/sql"
)

type VocabWord struct {
	ID          int
	Foreign1    string
	Foreign2    string
	Definitions string
	Level       string
}

type VocabWordRes struct {
	ID          int      `json:"id"`
	Foreign1    string   `json:"foreign1"`
	Foreign2    string   `json:"foreign2"`
	Definitions []string `json:"definitions"`
	Level       string   `json:"level"`
}

type VocabWordResByLevel struct {
	JLPTN1 []VocabWordRes `json:"jlptN1`
	JLPTN2 []VocabWordRes `json:"jlptN2`
	JLPTN3 []VocabWordRes `json:"jlptN3`
	JLPTN4 []VocabWordRes `json:"jlptN4`
	JLPTN5 []VocabWordRes `json:"jlptN5`
}

func (p *VocabWord) DbCreateWord(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO words(foreign1,foreign2,definitions,level) VALUES($1, $2, $3, $4) RETURNING id",
		p.Foreign1, p.Foreign2, p.Definitions, p.Level).Scan(&p.ID)

	if err != nil {
		return err
	}
	return nil
}
