package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/jedzeins/jlpt_api/dictionaryService/src/services"
)

func GetVocabWordsByLevel(w http.ResponseWriter, r *http.Request) {

	id := strings.TrimPrefix(r.URL.Path, "/vocab/level/")

	parsed, err := strconv.Atoi(id)
	if err != nil {
		w.Write([]byte("err in strconv"))
		return
	}

	if parsed < 1 || parsed > 5 {
		// return nil, errors.New("must be a full integer between 1 & 5")
		w.Write([]byte("must be a full integer between 1 & 5"))
		return
	}

	res, err := services.VocabWordsServiceByLevel(parsed, 11, 0)
	if err != nil {
		fmt.Errorf("err in vocab service: %w", err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
