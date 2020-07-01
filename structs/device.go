package structs

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Device struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty"`
	DeviceID            int                `json:"deviceId" bson:"deviceId"`
	DeviceName          string             `json:"deviceName" bson:"deviceName"`
	Battery             int                `json:"battery" bson:"battery"`
	DateTime            time.Time          `bson:"dateTime" json:"dateTime"`
	AirValue            int                `bson:"airValue" json:"airValue"`
	WaterValue          int                `bson:"waterValue" json:"waterValue"`
	SoilMoistureValue   int                `bson:"soilMoistureValue" json:"soilMoistureValue"`
	SoilMoisturePercent int                `bson:"soilMoisturePercent" json:"soilMoisturePercent"`
}

type DeviceData struct {
	UID                 string `bson:"uid" json:"uid"`
	DeviceID            int    `bson:"deviceId" json:"deviceId"`
	Battery             int    `json:"battery" bson:"battery"`
	AirValue            int    `bson:"airValue" json:"airValue"`
	WaterValue          int    `bson:"waterValue" json:"waterValue"`
	SoilMoistureValue   int    `bson:"soilMoistureValue" json:"soilMoistureValue"`
	SoilMoisturePercent int    `bson:"soilMoisturePercent" json:"soilMoisturePercent"`
	Token               string `bson:"token" json:"token"`
}
