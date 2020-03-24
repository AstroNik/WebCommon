package structs

import (
	"time"
)

type Sensor struct {
	//ID         primitive.ObjectID `bson:"_id,omitempty"`
	SensorId   int        `json:"sensorId" bson:"sensorId"`
	SensorName string     `json:"sensorName" bson:"SensorName"`
	DateTime   time.Time  `json:"dateTime" bson:"dateTime"`
	SensorData SensorData `json:"sensorData" bson:"sensorData"`
}

type SensorData struct {
	AirValue            int `json:"airValue" bson:"airValue"`
	WaterValue          int `json:"waterValue" bson:"waterValue"`
	SoilMoistureValue   int `json:"soilMoistureValue" bson:"soilMoistureValue"`
	SoilMoisturePercent int `json:"soilMoisturePercent" bson:"soilMoisturePercent"`
}
