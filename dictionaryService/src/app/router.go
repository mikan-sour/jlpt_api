package app

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jedzeins/jlpt_api/dictionaryService/src/controllers"
	"github.com/jedzeins/jlpt_api/dictionaryService/src/models"
)

const port = ":8080"

func makeRouter(app models.App) {

	http.HandleFunc("/health", controllers.HealthCheck)

	app.Router = &http.Server{
		Addr:           port,
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Printf("SERVING ON PORT %s", port)
	log.Fatal(app.Router.ListenAndServe())

}
