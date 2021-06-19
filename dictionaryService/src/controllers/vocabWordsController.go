package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jedzeins/jlpt_api/dictionaryService/src/models"
	"github.com/jedzeins/jlpt_api/dictionaryService/src/services"
)

func GetVocabWordsByLevel(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Can only send post requests"))
		return
	}

	decoder := json.NewDecoder(r.Body)
	var postRequest models.PostRequest
	err := decoder.Decode(&postRequest)
	if err != nil {
		decodeErr := errors.New(fmt.Errorf("error decoding the body: %w", err).Error())
		w.Write([]byte(decodeErr.Error()))
		return
	}

	if postRequest.Level < 0 || postRequest.Level > 5 {
		// return nil, errors.New("must be a full integer between 1 & 5")
		w.Write([]byte("must be a full integer between 1 & 5"))
		return
	}

	queryParams := r.URL.Query()

	var isByLevel bool
	if queryParams.Get("byLevel") != "" {
		isByLevel, err = strconv.ParseBool(queryParams.Get("byLevel"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("query param for `byLevel` must be `true` or `false`"))
			return
		}
	}

	var res interface{}

	if isByLevel {
		res, err = services.VocabWordsServiceGetWordsByLevel(postRequest.Text, postRequest.Limit, postRequest.Offset)

	} else {
		res, err = services.VocabWordsServiceGetWords(postRequest.Level, postRequest.Text, postRequest.Limit, postRequest.Offset)

	}
	doResponse(res, err, w, http.StatusOK)

}

func doResponse(res interface{}, err error, w http.ResponseWriter, code int) {
	if err != nil {
		fmt.Errorf("err in vocab service: %w", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("an error hath occured"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(res)
}
