package models

import "database/sql"

type VocabWord struct {
	ID          int
	Foreign1    string
	Foreign2    string
	Definitions string
	Level       string
}

type VocabWordRes struct {
	ID          int
	Foreign1    string
	Foreign2    string
	Definitions []string
	Level       string
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
