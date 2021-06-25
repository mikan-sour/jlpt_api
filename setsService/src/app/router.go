package app

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jedzeins/jlpt_api/setService/src/controllers"
	"github.com/jedzeins/jlpt_api/setService/src/models"
)

const port = ":8081"

func makeRouter(app models.App) {

	http.HandleFunc("/healthcheck", controllers.Healthcheck)

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
