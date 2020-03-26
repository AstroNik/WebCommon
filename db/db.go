package db

import (
	"context"
	"github.com/AstroNik/WebCommon/structs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func ConnectClient() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://devTeam:ecoders4@cluster0-grjmu.azure.mongodb.net/admin"))
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Connection Successful\n")
	return client
}

func InsertMoistureData(customerId string, sensor structs.Sensor) {
	log.Println("Sensor Data: ", sensor)
	client := ConnectClient()
	//eventually user will be passed in this function as an ID
	col := client.Database("234556314").Collection("SensorData")
	_, err := col.InsertOne(context.TODO(), sensor)
	if err != nil {
		log.Println("Cannot insert document")
	}

}

//func GetMoistureData() {
//	client := ConnectClient()
//
//}

func GetUserData() {
	//client := ConnectClient()
}
