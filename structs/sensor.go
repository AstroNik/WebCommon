package structs

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Sensor struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty"`
	DateTime            time.Time          `json:"dateTime" bson:"dateTime"`
	SensorName          string             `json:"sensorName" bson:"sensorName"`
	AirValue            int                `json:"airValue" bson:"airValue"`
	WaterValue          int                `json:"waterValue" bson:"waterValue"`
	SoilMoistureValue   int                `json:"soilMoistureValue" bson:"soilMoistureValue"`
	SoilMoisturePercent int                `json:"soilMoisturePercent" bson:"soilMoisturePercent"`
}

type SensorInfo struct {
	SensorName          string `json:"sensorName" bson:"sensorName"`
	AirValue            int    `json:"airValue" bson:"airValue"`
	WaterValue          int    `json:"waterValue" bson:"waterValue"`
	SoilMoistureValue   int    `json:"soilMoistureValue" bson:"soilMoistureValue"`
	SoilMoisturePercent int    `json:"soilMoisturePercent" bson:"soilMoisturePercent"`
}