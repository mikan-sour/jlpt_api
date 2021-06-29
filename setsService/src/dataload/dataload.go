package dataload

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/jedzeins/jlpt_api/setsService/src/database"
	"github.com/jedzeins/jlpt_api/setsService/src/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CheckIfDataExists() (bool, error) {
	var sets []models.Set

	cur, err := database.Collection.Find(database.Ctx, bson.D{})

	if err != nil {
		panic(err)
	}
	defer cur.Close(database.Ctx)

	if err := cur.All(database.Ctx, &sets); err != nil {
		panic(err)
	}

	if len(sets) == 0 {
		fmt.Println("No data, doing data load of sets")
		return false, nil
	}

	fmt.Println("Data already loaded")
	return true, nil
}

func DoDataload() error {

	data, err := ioutil.ReadFile("./src/dataload/dataload.json")
	if err != nil {
		fmt.Println(err)
		return err
	}

	var toBeLoaded []interface{}

	err = json.Unmarshal(data, &toBeLoaded)
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = database.Collection.InsertMany(database.Ctx, toBeLoaded)
	if err != nil {
		fmt.Println(err)
		return err
	}

	index := mongo.IndexModel{Keys: bson.M{"setName": "text"}, Options: nil}

	_, err = database.Collection.Indexes().CreateOne(context.Background(), index)
	if err != nil {
		panic(err)
	}

	return nil

}
