package structs

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Device struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty"`
	DeviceID            int                `bson:"deviceId" json:"deviceId"`
	Battery             int                `bson:"battery" json:"battery"`
	DateTime            time.Time          `bson:"dateTime" json:"dateTime"`
	AirValue            int                `bson:"airValue" json:"airValue"`
	WaterValue          int                `bson:"waterValue" json:"waterValue"`
	SoilMoistureValue   int                `bson:"soilMoistureValue" json:"soilMoistureValue"`
	SoilMoisturePercent int                `bson:"soilMoisturePercent" json:"soilMoisturePercent"`
}

type DeviceData struct {
	UID                 string `bson:"uid" json:"uid"`
	DeviceID            string `bson:"deviceId" json:"deviceId"`
	Battery             int    `bson:"battery" json:"battery"`
	AirValue            int    `bson:"airValue" json:"airValue"`
	WaterValue          int    `bson:"waterValue" json:"waterValue"`
	SoilMoistureValue   int    `bson:"soilMoistureValue" json:"soilMoistureValue"`
	SoilMoisturePercent int    `bson:"soilMoisturePercent" json:"soilMoisturePercent"`
}

type DSData struct {
	DeviceID            int       `bson:"deviceId" json:"deviceId"`
	DateTime            time.Time `bson:"dateTime" json:"dateTime"`
	SoilMoisturePercent int       `bson:"soilMoisturePercent" json:"soilMoisturePercent"`
}
