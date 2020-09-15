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

func Test() {
	log.Println("Connection To Client")
	client := ConnectClient()

	result, err := client.Database("Plant").ListCollectionNames(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	for _, coll := range result {
		log.Println(coll)
		log.Println("Hello")
	}
}
