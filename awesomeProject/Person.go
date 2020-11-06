package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Person struct {
	ID primitive.ObjectID `json:"_id, omitempty" bson:"_id,omitempty"`
	FirstName string `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname string `jsong:"lastname, omitempty" bson:"lastname,omitempty"`
}
