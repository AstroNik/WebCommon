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
	//TODO: replace pass/user with env var
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

	filter := options.Find()
	filter.SetSort(bson.D{{"_id", -1}})
	filter.SetLimit(1)

	cur, err := col.Find(context.TODO(), bson.D{}, filter)
	if err != nil {
		log.Println("Cannot retrieve document ERROR: ", err)
	}

	for cur.Next(context.TODO()) {
		err := cur.Decode(&sensorData)
		if err != nil {
			log.Println("Error decoding data ERROR: ", err)
		}
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	log.Printf("Document Found %+v\n", sensorData)
	return sensorData
}

func InsertUser(user structs.User) {
	client := ConnectClient()
	col := client.Database(user.UID).Collection(user.UID)
	_, err := col.InsertOne(context.TODO(), user)
	if err != nil {
		log.Println("Cannot insert document ERROR: ", err)
	}
}
