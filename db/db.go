package db

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func connectClient() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://devTeam:ecoders4@cluster0-grjmu.azure.mongodb.net/admin"))
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Connection Successful\n")
	return client
}

//func insertPlantData(sensor Sensor){
//	client := connectClient()
//
//}
