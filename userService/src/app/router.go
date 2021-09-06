package app

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jedzeins/jlpt_api/userService/src/controllers"
	"github.com/jedzeins/jlpt_api/userService/src/models"
)

const Port = ":8082"

func makeRouter(app models.App) {

	http.HandleFunc("/health", controllers.HealthCheck)
	http.HandleFunc("/signup", controllers.SignUp)

	app.Router = &http.Server{
		Addr:           Port,
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println("SERVING ON PORT " + Port)
	log.Fatal(app.Router.ListenAndServe())

}
