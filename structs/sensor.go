package structs

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Sensor struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	SensorData SensorData         `json:"sensorData" bson:"sensorData"`
}

type SensorData struct {
	DateTime            time.Time `json:"dateTime" bson:"dateTime"`
	SensorName          string    `json:"sensorName" bson:"sensorName"`
	AirValue            int       `json:"airValue" bson:"airValue"`
	WaterValue          int       `json:"waterValue" bson:"waterValue"`
	SoilMoistureValue   int       `json:"soilMoistureValue" bson:"soilMoistureValue"`
	SoilMoisturePercent int       `json:"soilMoisturePercent" bson:"soilMoisturePercent"`
}
