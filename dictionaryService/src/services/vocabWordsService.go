package services

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/jedzeins/jlpt_api/dictionaryService/src/models"
	"github.com/jedzeins/jlpt_api/dictionaryService/src/repository"
)

func VocabWordsService(level int, limit int, offset int) (*[]models.VocabWordRes, error) {
	if level < 1 || level > 5 {
		return nil, errors.New("must be a full integer between 1 & 5")
	}

	var actualLevel = fmt.Sprintf("JLPT N-%s", strconv.Itoa(level))

	return repository.DbGetWords(actualLevel, limit, offset)
}
