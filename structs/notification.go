package structs

/*
Code Written By
Nikhil Kapadia
991495131
*/

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Notification struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	NotificationID int                `json:"notificationId" bson:"notificationId"`
	DateTime       time.Time          `json:"dateTime" bson:"dateTime"`
	Title          string             `json:"title" bson:"title"`
	Content        string             `json:"content" bson:"content"`
	IsRead         bool               `json:"isRead" bson:"isRead"`
}
