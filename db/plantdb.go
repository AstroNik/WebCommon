package db

import (
	"context"
	"github.com/AstroNik/WebCommon/structs"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func GetPlantData(plantName string) structs.Plant {
	log.Printf("Finding Data on: %+v", plantName)
	client := ConnectClient()
	col := client.Database("Plant").Collection(plantName)

	filter := bson.D{
		{
			"commonName", plantName,
		},
	}

	var plantData structs.Plant

	_ = col.FindOne(context.TODO(), filter).Decode(&plantData)

	return plantData
}

func GetAllPlantData() []interface{} {
	var slice = make([]interface{}, 0)

	client := ConnectClient()
	col := client.Database("Plant").Collection("Plants")

	cur, err := col.Find(context.TODO(), bson.D{})

	if err != nil {
		log.Printf("error decoding GetAallPlantData %+v", err)
	}

	for cur.Next(context.TODO()) {
		var data structs.Plant

		err := cur.Decode(&data)
		if err != nil {
			log.Fatal(err)
		}
		slice = append(slice, data)
	}

	if err := cur.Err(); err != nil {
		log.Printf("error decoding GetAallPlantData %+v", err)
	}

	_ = cur.Close(context.TODO())

	log.Printf("Plant Data %+v", slice)

	return slice
}
