package dataload

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/jedzeins/jlpt_api/setService/src/database"
	"github.com/jedzeins/jlpt_api/setsService/src/models"
)

func DoDataload() error {

	data, err := ioutil.ReadFile("./src/dataload/dataload.json")
	if err != nil {
		fmt.Println(err)
		return err
	}
	// fmt.Println("data:  ", string(data))

	var toBeLoaded []models.Set

	err = json.Unmarshal(data, &toBeLoaded)
	if err != nil {
		fmt.Println(err)
		return err
	}
	// fmt.Printf("slice: %q\n", toBeLoaded)

	err = database.Collection.Insert(toBeLoaded)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil

}
