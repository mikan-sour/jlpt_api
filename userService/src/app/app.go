package app

import (
	"fmt"

	"github.com/jedzeins/jlpt_api/userService/src/database"
	"github.com/jedzeins/jlpt_api/userService/src/models"
)

var App = models.App{}

func StartApp() {
	err := database.Initialize()
	if err != nil {
		fmt.Errorf("error initializing the DB w/ data: %w", err)
		return
	}

	fmt.Println("DB CONNECTED ON PORT", database.Port)

	makeRouter(App)
}
