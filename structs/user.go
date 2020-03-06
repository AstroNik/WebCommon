package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	CustomerID int                `json:"customerId" bson:"customerId"`
	FirstName  string             `json:"firstName" bson:"firstName"`
	LastName   string             `json:"LastName" bson:"lastName"`
	Email      string             `json:"email" bson:"email"`
	DeviceId   []int              `json:"deviceId" bson:"deviceId"`
}

type Users struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Email    string             `json:"email" bson:"email"`
	UserID   string             `json:"userId" bson:"userId"`
	Password string             `json:"password" bson:"password"`
}
