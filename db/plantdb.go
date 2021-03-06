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
	"log"
)

func GetPlantData(plantName string) structs.Plant {
	log.Printf("Finding Data on: %+v", plantName)
	client := ConnectClient()
	col := client.Database("Plant").Collection("Plants")

	filter := bson.D{
		{"commonName", plantName},
	}

	var plantData structs.Plant

	_ = col.FindOne(context.TODO(), filter).Decode(&plantData)
	_ = client.Disconnect(context.TODO())
	return plantData
}

func GetAllPlantData() []structs.Plant {
	var plantData []structs.Plant

	client := ConnectClient()
	col := client.Database("Plant").Collection("Plants")

	cur, err := col.Find(context.TODO(), bson.D{})

	if err != nil {
		log.Printf("error decoding GetAllPlantData %+v", err)
	}

	for cur.Next(context.TODO()) {
		var data structs.Plant

		err := cur.Decode(&data)
		if err != nil {
			log.Fatal(err)
		}
		plantData = append(plantData, data)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	_ = cur.Close(context.TODO())

	log.Println(plantData)
	_ = client.Disconnect(context.TODO())
	return plantData
}
