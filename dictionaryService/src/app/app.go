package app

import (
	"fmt"

	"github.com/jedzeins/jlpt_api/dictionaryService/src/dataload"
	"github.com/jedzeins/jlpt_api/dictionaryService/src/models"
)

func StartApp() {
	app := models.App{}

	err := app.Initialize()
	if err != nil {
		fmt.Errorf("error initializing the DB w/ data: %w", err)
		return
	}

	if !app.CheckDB() {
		dataload.Dataload(app)
	} else {
		fmt.Println("TABLE EXISTS, SKIPPING DATALOAD")
	}

	makeRouter(app)

}
