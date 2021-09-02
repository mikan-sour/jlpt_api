package services

import (
	"log"
	"net/http"

	"github.com/jedzeins/jlpt_api/setsService/src/models"
	"github.com/jedzeins/jlpt_api/setsService/src/repositories"
	"github.com/jedzeins/jlpt_api/setsService/src/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func QuerySetsService(urlParams *models.SetRequestParamsUnParsed) (*models.SetResponse, *models.ApiError) {
	results, err := repositories.GetSets(urlParams)
	if err != nil {
		return nil, &models.ApiError{ErrorMessage: err.Error()}
	}

	return &models.SetResponse{StatusCode: http.StatusOK, Sets: *results}, nil
}

func QuerySetByIDService(id string) (*models.SetResponse, *models.ApiError) {
	result, err := repositories.GetSetById(id)
	if err != nil {
		return nil, &models.ApiError{ErrorMessage: err.Error()}
	}

	resSets := make([]models.Set, 0)
	resSets = append(resSets, *result)

	return &models.SetResponse{StatusCode: http.StatusOK, Sets: resSets}, nil
}

func DeleteSetById(id string) *models.ApiError {
	err := repositories.DeleteSetById(id)
	if err != nil {
		return &models.ApiError{ErrorMessage: err.Error()}
	}

	return nil
}

func PostNewSet(set models.Set) (*models.Set, *models.ApiError) {

	bsonSet, err := utils.SetToBson(&set)
	if err != nil {
		log.Fatal(err)
	}

	res, err := repositories.PostNewSet(bsonSet)
	if err != nil {
		return nil, &models.ApiError{ErrorMessage: err.Error()}
	}

	set.ID = *res

	return &set, nil
}

func PatchSetById(id string, set models.Set) (*models.Set, *models.ApiError) {

	bsonSet, err := utils.SetToBson(&set)
	if err != nil {
		log.Fatal(err)
	}

	err = repositories.UpdateSet(bsonSet, id)
	if err != nil {
		return nil, &models.ApiError{ErrorMessage: err.Error()}
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	set.ID = objID

	return &set, nil

}
