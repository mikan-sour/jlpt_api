package app

import (
	"fmt"

	"github.com/jedzeins/jlpt_api/dictionaryService/src/database"
	"github.com/jedzeins/jlpt_api/dictionaryService/src/dataload"
	"github.com/jedzeins/jlpt_api/dictionaryService/src/models"
)

var App = models.App{}

func StartApp() {

	err := database.Initialize()
	if err != nil {
		fmt.Errorf("error initializing the DB w/ data: %w", err)
		return
	}

	if !database.CheckDB() {
		dataload.Dataload()
	} else {
		fmt.Println("TABLE EXISTS, SKIPPING DATALOAD")
	}

	makeRouter(App)

}
