package db

import (
	"context"
	"github.com/AstroNik/WebCommon/structs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
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
	//eventually user will be passed in this function as an ID
	col := client.Database(customerId).Collection("SensorData")
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	_, err := col.InsertOne(ctx, sensor)
	if err != nil {
		log.Println("Cannot insert document ERROR: ", err)
	}

}

//func GetMoistureData() {
//	client := ConnectClient()
//
//}

func GetUserData() {
	//client := ConnectClient()
}
