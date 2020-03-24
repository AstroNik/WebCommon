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

func InsertPlantData(sensor structs.Sensor) {
	client := ConnectClient()
	//eventually user will be passed in this function as an ID
	col := client.Database("User").Collection("SensorData")
	_, err := col.InsertOne(context.TODO(), sensor)
	if err != nil {
		log.Println("Cannot insert document")
	}

}

func GetPlantData() {
	//client := ConnectClient()
}

func GetUserData() {
	//client := ConnectClient()
}
