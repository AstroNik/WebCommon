package structs

import (
	"time"
)

type Sensor struct {
	//ID         primitive.ObjectID `bson:"_id,omitempty"`
	SensorId            int       `bson:"sensorId" json:"sensorId"`
	SensorName          string    `bson:"SensorName" json:"sensorName"`
	DateTime            time.Time `bson:"dateTime" json:"dateTime"`
	AirValue            int       `bson:"airValue" json:"airValue"`
	WaterValue          int       `bson:"waterValue" json:"waterValue"`
	SoilMoistureValue   int       `bson:"soilMoistureValue" json:"soilMoistureValue"`
	SoilMoisturePercent int       `bson:"soilMoisturePercent" json:"soilMoisturePercent"`
}

type SensorData struct {
	AirValue            int `bson:"airValue" json:"airValue"`
	WaterValue          int `bson:"waterValue" json:"waterValue"`
	SoilMoistureValue   int `bson:"soilMoistureValue" json:"soilMoistureValue"`
	SoilMoisturePercent int `bson:"soilMoisturePercent" json:"soilMoisturePercent"`
}
