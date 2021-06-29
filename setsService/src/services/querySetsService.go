package services

import (
	"net/http"

	"github.com/jedzeins/jlpt_api/setsService/src/models"
	"github.com/jedzeins/jlpt_api/setsService/src/repositories"
)

func QuerySetsService(urlParams *models.SetRequestParamsUnParsed) (*models.SetResponse, *models.ApiError) {
	results, err := repositories.GetSets(urlParams)
	if err != nil {
		return nil, &models.ApiError{ErrorMessage: err.Error()}
	}

	return &models.SetResponse{StatusCode: http.StatusOK, Sets: *results}, nil
}
