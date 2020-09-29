package db

import (
	"context"
	"github.com/AstroNik/WebCommon/structs"
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
