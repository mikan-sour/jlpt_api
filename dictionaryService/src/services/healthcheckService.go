package services

import (
	"time"

	"github.com/jedzeins/jlpt_api/dictionaryService/src/models"
)

func HealthcheckService() models.HealthCheck {
	t := time.Now().Unix()
	tm := time.Unix(t, 0)

	return models.HealthCheck{
		Status: "Dictionary Service is Good! 順調している！",
		Time:   tm,
	}

}
