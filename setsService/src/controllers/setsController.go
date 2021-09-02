package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/jedzeins/jlpt_api/setsService/src/models"
	"github.com/jedzeins/jlpt_api/setsService/src/services"
	"github.com/jedzeins/jlpt_api/setsService/src/utils"
)

func HandleSets(w http.ResponseWriter, r *http.Request) {

	urlParams, parseURLErr := utils.ParseSetUrlCheckStrings(r.URL)
	if parseURLErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(parseURLErr.ErrorMessage))
		return
	}

	queryHasId := false

	if urlParams.Id != "" {
		queryHasId = true
	}

	switch meth := r.Method; meth {

	case "POST":
		postNewSet(w, r)
	case "GET":
		if queryHasId {
			querySetById(w, r, urlParams.Id)
		} else {
			querySets(w, r, urlParams)
		}
	case "DELETE":
		deleteSetById(w, r, urlParams.Id)
	case "PATCH":
		updateSetById(w, r, urlParams.Id)
	default:
		fmt.Println(r.Method)
	}
}

func querySets(w http.ResponseWriter, r *http.Request, urlParams *models.SetRequestParamsUnParsed) {

	response, err := services.QuerySetsService(urlParams)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func querySetById(w http.ResponseWriter, r *http.Request, id string) {

	response, err := services.QuerySetByIDService(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func postNewSet(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var postRequest models.Set
	err := decoder.Decode(&postRequest)
	if err != nil {
		decodeErr := errors.New(fmt.Errorf("error decoding the body: %w", err).Error())
		w.Write([]byte(decodeErr.Error()))
		return
	}

	response, ApiError := services.PostNewSet(postRequest)
	if ApiError != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ApiError)
		return
	}

	fmt.Println(response)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func updateSetById(w http.ResponseWriter, r *http.Request, id string) {

	decoder := json.NewDecoder(r.Body)
	var postRequest models.Set
	err := decoder.Decode(&postRequest)
	if err != nil {
		decodeErr := errors.New(fmt.Errorf("error decoding the body: %w", err).Error())
		w.Write([]byte(decodeErr.Error()))
		return
	}

	response, ApiError := services.PatchSetById(id, postRequest)
	if ApiError != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ApiError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func deleteSetById(w http.ResponseWriter, r *http.Request, id string) {
	err := services.DeleteSetById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("deleted"))
}
