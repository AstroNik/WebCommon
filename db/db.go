package db

import (
	"context"
	"github.com/AstroNik/WebCommon/structs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func ConnectClient() *mongo.Client {
	clientOption := options.Client().ApplyURI("mongodb+srv://devTeam:ecoders4@cluster0-grjmu.azure.mongodb.net/admin")
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Println("mongo.Connect() ERROR: ", err)
	}
	return client
}

func InsertMoistureData(customerId string, sensor structs.Sensor) {
	log.Println("Sensor Data: ", sensor)
	client := ConnectClient()
	col := client.Database(customerId).Collection("SensorData")
	_, err := col.InsertOne(context.TODO(), sensor)
	if err != nil {
		log.Println("Cannot insert document ERROR: ", err)
	}

}

func GetMoistureData(customerId string) structs.Sensor {
	sensorData := structs.Sensor{}
	client := ConnectClient()
	col := client.Database(customerId).Collection("SensorData")

	filter := bson.M{
		"$sort": bson.M{
			"datetime": -1,
		},
	}

	cur, err := col.Find(context.TODO(), filter)
	if err != nil {
		log.Println("Cannot retrieve document ERROR: ", err)
	}
	for cur.Next(context.TODO()) {
		cur.Decode(&sensorData)
	}
	log.Printf("Document Found %+v\n", sensorData)
	return sensorData
}

func GetUserData() {
	//client := ConnectClient()
}
