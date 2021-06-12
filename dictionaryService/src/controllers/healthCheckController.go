package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jedzeins/jlpt_api/dictionaryService/src/services/healthcheckService"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	status := healthcheckService.HealthcheckService()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
	fmt.Println("API ROUTER MOUNTED")
}
