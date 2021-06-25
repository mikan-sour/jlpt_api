package main

import (
	"log"

	"github.com/jedzeins/jlpt_api/setsService/src/database"
	"github.com/jedzeins/jlpt_api/setsService/src/dataload"
)

func main() {
	err := database.ConnectMongo()
	if err != nil {
		log.Fatal("error connecting to MongoDB: %w", err)
	}

	err = dataload.DoDataload()

	if err != nil {
		log.Fatal("error in dataload into MongoDB: %w", err)
	}

}
