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

	http.HandleFunc("/sets", controllers.HandleSets)
	http.HandleFunc("/health", controllers.HealthCheck)
	// http.HandleFunc("/sets", controllers.QuerySets)
	// http.HandleFunc("/setsById", controllers.QuerySetById)
	// http.HandleFunc("/setsNew", controllers.PostNewSet)

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
