package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TO BE - setWords will be a more complex object.
// type SetWord struct {
// 	ID int `json:"id,omitempty" bson:"id,omitempty"`
// }

type Set struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Owner    string             `json:"owner" bson:"owner"`
	SetName  string             `json:"setName" bson:"setName"`
	SetWords []int              `json:"setWords" bson:"setWords"`
	IsPublic bool               `json:"isPublic" bson:"isPublic"`
}
