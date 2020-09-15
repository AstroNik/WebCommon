package db

import (
	"context"
	"github.com/AstroNik/WebCommon/structs"
	"go.mongodb.org/mongo-driver/bson"
	"log"
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
