package main

import (
	"fmt"

	"github.com/jedzeins/jlpt_api/userService/src/app"
)

func main() {
	app.StartApp()
	fmt.Printf("STARTING USER-SERVICE ON PORT %s\n", ":8082")
}
