package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UID       string             `json:"uid" bson:"uid"`
	FirstName string             `json:"firstName" bson:"firstName"`
	LastName  string             `json:"LastName" bson:"lastName"`
	Email     string             `json:"email" bson:"email"`
	DeviceId  []int              `json:"deviceId" bson:"deviceId"`
	Token     string             `json:"token" bson:"token"`
}

type NewUser struct {
	UID       string `json:"uid" bson:"uid"`
	Email     string `json:"email" bson:"email"`
	FirstName string `json:"firstName" bson:"firstName"`
	LastName  string `json:"LastName" bson:"lastName"`
}
