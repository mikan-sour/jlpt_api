package app

import (
	"log"

	"github.com/jedzeins/jlpt_api/setsService/src/database"
	"github.com/jedzeins/jlpt_api/setsService/src/dataload"
	"github.com/jedzeins/jlpt_api/setsService/src/models"
)

var App = models.App{}

func StartApp() {
	err := database.ConnectMongo()
	if err != nil {
		log.Fatal("error connecting to MongoDB: %w", err)
	}

	check, err := dataload.CheckIfDataExists()
	if err != nil {
		log.Fatal("error in checking if data was loaded already: %w", err)
	}

	if !check {
		err = dataload.DoDataload()
		if err != nil {
			log.Fatal("error in dataload into MongoDB: %w", err)
		}
	}

	makeRouter(App)
}
