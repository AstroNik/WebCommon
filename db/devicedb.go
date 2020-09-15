package db

import (
	"context"
	"fmt"
	"github.com/AstroNik/WebCommon/structs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

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

	options := options.FindOne()
	options.SetSort(bson.D{{"_id", -1}})

	deviceIds := GetUniqueDevices(uid)

	for i := range deviceIds {
		tempData := structs.Device{}
		_ = col.FindOne(context.TODO(), bson.D{{"deviceId", deviceIds[i]}}, options).Decode(&tempData)
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

func GetAllMoistureData(uid string, timezone string) []interface{} {
	client := ConnectClient()
	col := client.Database(uid).Collection("Device")

	deviceIds := GetUniqueDevices(uid)

	var slice = make([]interface{}, 0)

	for i := range deviceIds {
		var allData []structs.DSData

		opts := options.Find().SetProjection(bson.D{
			{"_id", 0},
			{"deviceId", 1},
			{"dateTime", 1},
			{"soilMoisturePercent", 1},
		})

		loc, _ := time.LoadLocation(timezone)

		filter := bson.D{
			{"deviceId", deviceIds[i]},
			{"dateTime", bson.M{
				"$gte": DateBeginning(time.Now(), loc),
				"$lt":  DateEnd(time.Now(), loc),
			}},
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
			allData = append(allData, data)
		}

		if err := cur.Err(); err != nil {
			log.Fatal(err)
		}

		_ = cur.Close(context.TODO())
		slice = append(slice, allData)
	}

	fmt.Printf("Found multiple documents: %+v\n", slice)

	return slice
}

func GetSpecificDayChartData(uid string, deviceId int, date time.Time, timezone string) []structs.DSData {
	client := ConnectClient()
	col := client.Database(uid).Collection("Device")

	var deviceData []structs.DSData

	opts := options.Find().SetProjection(bson.D{
		{"_id", 0},
		{"deviceId", 1},
		{"dateTime", 1},
		{"soilMoisturePercent", 1},
	})

	loc, _ := time.LoadLocation(timezone)

	filter := bson.D{
		{"deviceId", deviceId},
		{"dateTime", bson.M{
			"$gte": DateBeginning(date, loc),
			"$lt":  DateEnd(date, loc),
		}},
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
		deviceData = append(deviceData, data)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	_ = cur.Close(context.TODO())

	fmt.Printf("Found multiple documents: %+v\n", deviceData)

	return deviceData
}
