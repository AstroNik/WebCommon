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
	"log"
)

func PushNotification(uid string, notification structs.Notification) {
	client := ConnectClient()

	col := client.Database(uid).Collection("Notifications")
	_, err := col.InsertOne(context.TODO(), notification)
	if err != nil {
		log.Println("Cannot insert document ERROR: ", err)
	}
	_ = client.Disconnect(context.TODO())
}

func GetNotifications(uid string) []structs.Notification {
	client := ConnectClient()
	col := client.Database(uid).Collection("Notifications")

	cur, err := col.Find(context.TODO(), bson.D{{"isRead", false}})

	if err != nil {
		log.Fatal(err)
	}

	var notifs []structs.Notification

	for cur.Next(context.TODO()) {
		var data structs.Notification

		err := cur.Decode(&data)
		if err != nil {
			log.Fatal(err)
		}
		notifs = append(notifs, data)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	_ = cur.Close(context.TODO())

	fmt.Printf("Found multiple documents: %+v\n", notifs)
	_ = client.Disconnect(context.TODO())
	return notifs
}
