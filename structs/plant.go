package structs

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Plant struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	BotanicalName string             `json:"botanicalName" bson:"botanicalName"` //Ex. Epipremnum aureum
	CommonName    string             `json:"commonName" bson:"commonName"`       //Ex. Golden pothos, silver vine, taro vine
	PlantType     string             `json:"plantType" bson:"plantType"`         //Ex. Trailing Vine
	MatureSize    string             `json:"matureSize" bson:"matureSize"`       // Facts about plants
	Care          Care               `json:"growthInfo" bson:"growthInfo"`
}

type Care struct {
	LightLevel string `json:"lightLevel" bson:"lightLevel"` //Sun Exposure. Indirect / Direct
	Humidity   string `json:"humidity" bson:"humidity"`     //Room humidity
	Duration   string `json:"duration" bson:"duration"`     //How often
	Direction  string `json:"direction" bson:"direction"`   //Near roots or far?
}
