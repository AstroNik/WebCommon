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

func GetMoistureData(uid string) []structs.Device {
	var deviceData []structs.Device
	client := ConnectClient()
	col := client.Database(uid).Collection("Device")

	filter := options.FindOne()
	filter.SetSort(bson.D{{"_id", -1}})

	deviceIds := GetUniqueDevices(uid)

	for i := range deviceIds {
		tempData := structs.Device{}
		_ = col.FindOne(context.TODO(), bson.D{{"deviceId", deviceIds[i]}}, filter).Decode(&tempData)
		deviceData = append(deviceData, tempData)
	}

	log.Printf("Document Found %+v\n", deviceData)
	return deviceData
}

func GetUniqueDevices(uid string) []int32 {
	client := ConnectClient()
	col := client.Database(uid).Collection("Device")

	deviceIds, err := col.Distinct(context.TODO(), "deviceId", bson.D{{}})

	if err != nil {
		log.Println("Error decoding data ERROR: ", err)
	}

	convertedIds := make([]int32, len(deviceIds))

	for i := range deviceIds {
		convertedIds[i] = deviceIds[i].(int32)
	}

	return convertedIds
}

func GetAllMoistureData(uid string) [][]structs.DSData {
	client := ConnectClient()
	col := client.Database(uid).Collection("Device")

	deviceIds := GetUniqueDevices(uid)

	var allData [][]structs.DSData

	for i := range deviceIds {
		opts := options.Find().SetProjection(bson.D{
			{"_id", 0},
			{"deviceId", 1},
			{"dateTime", 1},
			{"soilMoisturePercent", 1},
		})

		filter := bson.D{
			{"deviceId", deviceIds[i]},
		}

		cur, err := col.Find(context.TODO(), filter, opts)

		if err != nil {
			log.Fatal(err)
		}

		for cur.Next(context.TODO()) {
			var data structs.DSData

			err := cur.Decode(&data)
			if err != nil {
				log.Fatal(err)
			}
			allData[i] = append(allData[i], data)
		}

		if err := cur.Err(); err != nil {
			log.Fatal(err)
		}

		_ = cur.Close(context.TODO())
	}

	fmt.Printf("Found multiple documents: %+v\n", allData)

	return allData
}

func InsertUser(user structs.NewUser) {
	client := ConnectClient()
	col := client.Database(user.UID).Collection("User")
	_, err := col.InsertOne(context.TODO(), user)
	if err != nil {
		log.Println("Cannot insert document ERROR: ", err)
	}
}
