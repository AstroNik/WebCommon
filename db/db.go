package db

import (
	"context"
	"fmt"
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

func InsertMoistureData(uid string, sensor structs.Device) {
	log.Println(uid)
	log.Println("Device Data: ", sensor)
	client := ConnectClient()
	col := client.Database(uid).Collection("Device")
	_, err := col.InsertOne(context.TODO(), sensor)
	if err != nil {
		log.Println("Cannot insert document ERROR: ", err)
	}
}

func GetMoistureData(uid string) structs.Device {
	deviceData := structs.Device{}
	client := ConnectClient()
	col := client.Database(uid).Collection("Device")

	filter := options.Find()
	filter.SetSort(bson.D{{"_id", -1}})
	filter.SetLimit(1)

	cur, err := col.Find(context.TODO(), bson.D{}, filter)
	if err != nil {
		log.Println("Cannot retrieve document ERROR: ", err)
	}

	for cur.Next(context.TODO()) {
		err := cur.Decode(&deviceData)
		if err != nil {
			log.Println("Error decoding data ERROR: ", err)
		}
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	_ = cur.Close(context.TODO())

	log.Printf("Document Found %+v\n", deviceData)
	return deviceData
}

func GetUnqiueDevices(uid string) []int {
	var deviceIds []int
	client := ConnectClient()
	col := client.Database(uid).Collection("Device")

	fmt.Println(col.Distinct(context.TODO(), "deviceId", &deviceIds))

	return deviceIds
}

func InsertUser(user structs.NewUser) {
	client := ConnectClient()
	col := client.Database(user.UID).Collection("User")
	_, err := col.InsertOne(context.TODO(), user)
	if err != nil {
		log.Println("Cannot insert document ERROR: ", err)
	}
}
