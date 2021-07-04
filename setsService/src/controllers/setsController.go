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

func QuerySets(w http.ResponseWriter, r *http.Request) {

	// if r.Method != "POST" {
	// 	w.WriteHeader(http.StatusMethodNotAllowed)
	// 	w.Write([]byte("Can only send post requests"))
	// 	return
	// }

	urlParams, parseURLErr := utils.ParseSetUrlCheckStrings(r.URL)
	if parseURLErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(parseURLErr.ErrorMessage))
		return
	}

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

func QuerySetById(w http.ResponseWriter, r *http.Request) {
	urlParams, parseURLErr := utils.ParseSetUrlId(r.URL)
	if parseURLErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(parseURLErr.ErrorMessage))
		return
	}

	if r.Method == "DELETE" {
		err := services.DeleteSetById(urlParams.Id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("deleted"))
		return
	}

	if r.Method == "GET" {
		response, err := services.QuerySetByIDService(urlParams.Id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

}

func PostNewSet(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Can only send post requests"))
		return
	}

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
	return

}
