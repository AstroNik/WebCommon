package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type Plant struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	PlantName   string             `json:"plantName" bson:"plantName"`
	Description string             `json:"description" bson:"description"`
	GrowthInfo  string             `json:"growthInfo" bson:"growthInfo"`
}
