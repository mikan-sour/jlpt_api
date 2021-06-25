package services

import (
	"time"

	"github.com/jedzeins/jlpt_api/setsService/src/models"
)

func HealthcheckService() models.HealthCheck {
	t := time.Now().Unix()
	tm := time.Unix(t, 0)

	return models.HealthCheck{
		Status: "Sets Service is Good! 順調している！",
		Time:   tm,
	}

}
