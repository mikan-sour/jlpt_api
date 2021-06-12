package repository

import (
	"fmt"
	"strings"

	"github.com/jedzeins/jlpt_api/dictionaryService/src/database"
	"github.com/jedzeins/jlpt_api/dictionaryService/src/models"
)

func DbGetWords(level string, limit int, offset int) (*[]models.VocabWordRes, error) {

	rows, err := database.DB.Query(
		"SELECT * FROM words WHERE level = $1 ORDER BY id LIMIT $2 OFFSET $3",
		level, limit, offset)

	// rows, err := database.DB.Query("SELECT * FROM words WHERE level = 'JLPT N-4' ORDER BY id LIMIT 10 OFFSET 1")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	fmt.Println(rows.Next())

	words := []models.VocabWordRes{}

	for rows.Next() {
		var (
			p      models.VocabWordRes
			defs   string
			holder []string
		)

		if err := rows.Scan(&p.ID, &p.Foreign1, &p.Foreign2, &defs, &p.Level); err != nil {
			return nil, err
		}

		if strings.Contains(defs, ";") {
			holder = strings.Split(defs, ";")
			for _, i := range holder {
				p.Definitions = append(p.Definitions, strings.TrimSpace(i))
			}
		} else {
			p.Definitions = append(p.Definitions, defs)
		}

		words = append(words, p)
	}

	return &words, nil
}
