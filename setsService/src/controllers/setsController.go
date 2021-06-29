package controllers

import (
	"encoding/json"
	"net/http"

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
	}

	// var sets []models.Set
	// response := models.SetResponse{
	// 	StatusCode: http.StatusOK,
	// 	Sets:       sets,
	// }

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
