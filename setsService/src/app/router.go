package app

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jedzeins/jlpt_api/setsService/src/controllers"
	"github.com/jedzeins/jlpt_api/setsService/src/models"
)

const port = ":8081"

func makeRouter(app models.App) {

	http.HandleFunc("/healthcheck", controllers.HealthCheck)
	http.HandleFunc("/sets", controllers.QuerySets)

	app.Router = &http.Server{
		Addr:           port,
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println("SERVING ON PORT " + port)
	log.Fatal(app.Router.ListenAndServe())

}
