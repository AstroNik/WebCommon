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

func GetAllPlantData() []structs.Plant {

	return []structs.Plant{}
}
