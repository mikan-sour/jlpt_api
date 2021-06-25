package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TO BE - setWords will be a more complex object.
// type SetWord struct {
// 	ID int `json:"id,omitempty" bson:"id,omitempty"`
// }

type Set struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Owner    string             `json:"owner,omitempty" bson:"owner,omitempty"`
	SetName  string             `json:"setName,omitempty" bson:"setName,omitempty"`
	SetWords []int              `json:"setWords,omitempty" bson:"setWords,omitempty"`
}
