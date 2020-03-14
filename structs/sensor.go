package structs

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Sensor struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	SensorName string             `json:"sensorName" bson:"sensorName"`
	DateTime   string             `json:"dateTime" bson:"dateTime"`
	SensorData SensorData         `json:"sensorData" bson:"sensorData"`
}

type SensorData struct {
	AirValue            int `json:"airValue" bson:"airValue"`
	WaterValue          int `json:"waterValue" bson:"waterValue"`
	SoilMoistureValue   int `json:"soilMoistureValue" bson:"soilMoistureValue"`
	SoilMoisturePercent int `json:"soilMoisturePercent" bson:"soilMoisturePercent"`
}
