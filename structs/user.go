package structs

/*
Code Written By
Nikhil Kapadia
991495131
*/

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NewUser struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UID       string             `json:"uid" bson:"uid"`
	Email     string             `json:"email" bson:"email"`
	FirstName string             `json:"firstName" bson:"firstName"`
	LastName  string             `json:"LastName" bson:"lastName"`
	Token     string             `json:"token" bson:"token"`
	Devices   interface{}        `json:"devices" bson:"devices"`
}

type User struct {
	UID string `json:"uid" bson:"uid"`
}

type UserRetrieval struct {
	UID       string      `json:"uid" bson:"uid"`
	Email     string      `json:"email" bson:"email"`
	FirstName string      `json:"firstName" bson:"firstName"`
	LastName  string      `json:"LastName" bson:"lastName"`
	Devices   interface{} `json:"devices" bson:"devices"`
}
