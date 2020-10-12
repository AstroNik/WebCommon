package db

/*
Code Written By
Nikhil Kapadia
991495131
*/

import (
	"context"
	"fmt"
	"github.com/AstroNik/WebCommon/structs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func PushNotification(uid string, notification structs.Notification) {
	client := ConnectClient()

	col := client.Database(uid).Collection("Notifications")
	_, err := col.InsertOne(context.TODO(), notification)
	if err != nil {
		log.Println("Cannot insert Notification ERROR: ", err)
	}
	_ = client.Disconnect(context.TODO())
}

func GetNotifications(uid string) []structs.Notification {
	client := ConnectClient()
	col := client.Database(uid).Collection("Notifications")

	option := options.FindOne()
	option.SetSort(bson.D{{"_id", -1}})

	deviceIds := GetUniqueDevices(uid)
	var notifs []structs.Notification

	for i := range deviceIds {
		filter := bson.D{
			{"deviceId", deviceIds[i]},
			{"isRead", false},
		}

		tempData := structs.Notification{}
		_ = col.FindOne(context.TODO(), filter, option).Decode(&tempData)
		if tempData.DeviceID == 0 {

		} else {
			notifs = append(notifs, tempData)
		}

	}

	fmt.Printf("Found multiple Notification Documents: %+v\n", notifs)
	_ = client.Disconnect(context.TODO())
	return notifs
}

func UpdateNotification(uid string, notifId int) {
	client := ConnectClient()
	col := client.Database(uid).Collection("Notifications")

	filter := bson.M{"notificationId": notifId}

	update := bson.M{"$set": bson.M{"isRead": true}}

	_ = col.FindOneAndUpdate(context.TODO(), filter, update)

	log.Print("Notification Updated")
	_ = client.Disconnect(context.TODO())
}
