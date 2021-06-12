package app

import (
	"fmt"

	"github.com/jedzeins/jlpt_api/src/dataload"
	"github.com/jedzeins/jlpt_api/src/models"
)

func StartApp() {
	app := models.App{}

	err := app.Initialize()
	if err != nil {
		fmt.Errorf("error initializing the DB w/ data: %w", err)
		return
	}

	dataload.Dataload(&app)

	fmt.Println("we did it!")

}
