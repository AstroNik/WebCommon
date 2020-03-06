package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type Device struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	DeviceID   int                `json:"deviceId" bson:"deviceId"`
	AirValue   int                `json:"airValue" bson:"airValue"`
	WaterValue int                `json:"waterValue" bson:"waterValue"`
}
