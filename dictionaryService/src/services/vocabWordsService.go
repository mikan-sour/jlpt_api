package services

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jedzeins/jlpt_api/dictionaryService/src/models"
	"github.com/jedzeins/jlpt_api/dictionaryService/src/repository"
)

func VocabWordsServiceGetWords(level int, text string, limit int, offset int) (*[]models.VocabWordRes, error) {

	var actualLevel = fmt.Sprintf("JLPT N-%s", strconv.Itoa(level))
	rows, err := repository.DoQuery(actualLevel, text, limit, offset)
	if err != nil {
		panic(err) // handle this!
	}
	defer rows.Close()

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

func VocabWordsServiceGetWordsByLevel(text string, limit int, offset int) (*models.VocabWordResByLevel, error) {
	rows, err := repository.DoQuery(repository.LevelPlaceholder, text, limit, offset)
	if err != nil {
		panic(err) // handle this!
	}

	defer rows.Close()

	words := models.VocabWordResByLevel{}

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

		switch p.Level {
		case "JLPT N-1":
			words.JLPTN1 = append(words.JLPTN1, p)
		case "JLPT N-2":
			words.JLPTN2 = append(words.JLPTN2, p)
		case "JLPT N-3":
			words.JLPTN3 = append(words.JLPTN3, p)
		case "JLPT N-4":
			words.JLPTN4 = append(words.JLPTN4, p)
		case "JLPT N-5":
			words.JLPTN5 = append(words.JLPTN5, p)
		default:
			fmt.Println(p.Level)
		}

	}

	return &words, nil

}
