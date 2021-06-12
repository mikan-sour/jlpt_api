package healthcheckService

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jedzeins/jlpt_api/src/models"
)

func healthcheckService(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Unix()
	tm := time.Unix(t, 0)

	var status = models.HealthCheck{
		Status: "Good! 順調している！",
		Time:   tm,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
	fmt.Println("API ROUTER MOUNTED")
}
