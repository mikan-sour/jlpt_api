package repositories

import (
	"log"
	"reflect"

	"github.com/jedzeins/jlpt_api/setsService/src/database"
	"github.com/jedzeins/jlpt_api/setsService/src/models"
	"github.com/jedzeins/jlpt_api/setsService/src/utils"

	// "go.mongodb.org/mongo-driver/bson"
	"gopkg.in/mgo.v2/bson"
)

func makeBSONQuery(urlParams *models.SetRequestParamsUnParsed) (bson.M, error) {
	queryObject := bson.M{}

	// use reflect to iterate through the urlParams and add 🔑 and value if value is not blank
	v := reflect.ValueOf(*urlParams)
	typeOfURLParams := v.Type()
	for i := 0; i < v.NumField(); i++ {

		name := utils.LowerCaseString(typeOfURLParams.Field(i).Name)
		val := v.Field(i).Interface()

		if val != "" {
			if name == "offset" || name == "limit" || name == "isMine" { // must figure out how to query w/ Offsets & such
				continue
			}
			if name == "isPublic" {
				boolVal, err := utils.ParseStringToBool(val)
				if err != nil {
					log.Fatal(err)
				}
				queryObject[name] = boolVal
			}
			if name == "setName" {
				queryObject["$text"] = bson.M{"$search": val}
			} else {
				continue // catch all.
			}

		}
	}

	return queryObject, nil
}

func GetSets(urlParams *models.SetRequestParamsUnParsed) (*[]models.Set, error) {

	query, _ := makeBSONQuery(urlParams)

	// fakeQuery := bson.M{
	// 	"$text":    bson.M{"$search": urlParams.SetName},
	// 	"isPublic": true,
	// }

	var responseSets = []models.Set{}

	cur, err := database.Collection.Find(database.Ctx, query)

	if err != nil {
		return nil, err
	}
	defer cur.Close(database.Ctx)

	for cur.Next(database.Ctx) {
		result := models.Set{}
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}

		responseSets = append(responseSets, result)

	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return &responseSets, nil
}
