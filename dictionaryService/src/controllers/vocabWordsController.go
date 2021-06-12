package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jedzeins/jlpt_api/dictionaryService/src/services"
)

func GetVocabWords(w http.ResponseWriter, r *http.Request) {
	res, err := services.VocabWordsService(2, 10, 0)
	if err != nil {
		fmt.Errorf("err in vocab service: %w", err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
