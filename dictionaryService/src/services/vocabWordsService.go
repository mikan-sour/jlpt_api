package services

import (
	"fmt"
	"strconv"

	"github.com/jedzeins/jlpt_api/dictionaryService/src/models"
	"github.com/jedzeins/jlpt_api/dictionaryService/src/repository"
)

func VocabWordsServiceByLevel(level int, limit int, offset int) (*[]models.VocabWordRes, error) {

	var actualLevel = fmt.Sprintf("JLPT N-%s", strconv.Itoa(level))

	return repository.DbGetWordsByLevel(actualLevel, limit, offset)
}
