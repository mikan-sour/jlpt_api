package repository

import (
	"database/sql"
	"fmt"

	"github.com/jedzeins/jlpt_api/dictionaryService/src/database"
)

var LevelPlaceholder = "JLPT N-0"

func makeQueryString(level string, text string, limit int, offset int) (string, error) { // Need to make an error...
	var queryStringAltogether string
	queryStringRoot := fmt.Sprint("SELECT * FROM words WHERE 1=1")
	queryStringLevel := fmt.Sprintf(" AND level = '%s'", level)
	queryStringText := fmt.Sprintf(` AND(
		to_tsvector(foreign1) @@ to_tsquery('%s') OR
		to_tsvector(foreign2) @@ to_tsquery('%s') OR
		to_tsvector(definitions) @@ to_tsquery('%s')
	)`, text, text, text)

	queryStringLimitOffset := fmt.Sprintf(" ORDER BY id LIMIT %v OFFSET %v", limit, offset)

	queryStringAltogether += queryStringRoot

	if level != "" && level != LevelPlaceholder {
		queryStringAltogether += queryStringLevel
	}

	if text != "" {
		queryStringAltogether += queryStringText
	}

	queryStringAltogether += queryStringLimitOffset

	return queryStringAltogether, nil

}

func DoQuery(level string, text string, limit int, offset int) (*sql.Rows, error) {
	query, err := makeQueryString(level, text, limit, offset)
	if err != nil {
		panic(err) // handle this!
	}

	rows, err := database.DB.Query(query)

	if err != nil {
		return nil, err
	}

	return rows, err
}
