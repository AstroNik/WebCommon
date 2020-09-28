package db

import (
	"context"
	"github.com/AstroNik/WebCommon/structs"
	"go.mongodb.org/mongo-driver/bson"
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
}

func RetrieveUserInfo(uid string) structs.UserRetrieval {
	var user structs.UserRetrieval
	client := ConnectClient()
	col := client.Database(uid).Collection("User")
	_ = col.FindOne(context.TODO(), bson.D{}).Decode(&user)
	log.Print(user)
	return user
}

func AddDeviceToProfile(uid string, deviceId int, deviceName string) {
	client := ConnectClient()
	col := client.Database(uid).Collection("User")

	idString := strconv.Itoa(deviceId)

	filter := bson.M{"uid": uid}

	update := bson.M{
		"$set": bson.M{"Devices": bson.M{idString: deviceName}},
	}
	//option := options.FindOneAndUpdate()

	err := col.FindOneAndUpdate(context.TODO(), filter, update)
	if err != nil {
		log.Println("Cannot insert document ERROR: ", err)
	}
}
