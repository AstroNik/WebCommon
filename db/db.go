package db

/*
Code Written By
Nikhil Kapadia
991495131
*/

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func ConnectClient() *mongo.Client {
	//TODO: replace pass/user with env var
	clientOption := options.Client().ApplyURI("mongodb+srv://username:pass@cluster0-grjmu.azure.mongodb.net/admin")
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Println("mongo.Connect() ERROR: ", err)
	}
	return client
}

func DateBeginning(t time.Time, location *time.Location) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, location)
}

func DateEnd(t time.Time, location *time.Location) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 23, 59, 59, 999, location)
}
