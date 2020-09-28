package db

/*
Code Written By
Nikhil Kapadia
991495131
*/

import (
	"context"
	"github.com/AstroNik/WebCommon/structs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strconv"
)

func InsertUser(user structs.NewUser) {
	client := ConnectClient()
	col := client.Database(user.UID).Collection("User")
	_, err := col.InsertOne(context.TODO(), user)
	if err != nil {
		log.Println("Cannot insert document ERROR: ", err)
	}
	_ = client.Disconnect(context.TODO())
}

func RetrieveUserInfo(uid string) structs.UserRetrieval {
	var user structs.UserRetrieval
	client := ConnectClient()
	col := client.Database(uid).Collection("User")
	_ = col.FindOne(context.TODO(), bson.D{}).Decode(&user)
	_ = client.Disconnect(context.TODO())
	log.Print(user)
	return user
}

func AddDeviceToProfile(uid string, deviceId int, deviceName string) {
	client := ConnectClient()
	col := client.Database(uid).Collection("User")

	idString := strconv.Itoa(deviceId)

	filter := bson.M{"uid": uid}

	update := bson.M{"$set": bson.M{"devices." + idString: deviceName}}

	option := options.FindOneAndUpdate()

	_ = col.FindOneAndUpdate(context.TODO(), filter, update, option)

	log.Print("Device Added")
	_ = client.Disconnect(context.TODO())
}
