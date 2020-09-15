package structs

type FCMMessage struct {
	To           string      `json:"to" bson:"to"`
	Registration []string    `json:"registration _ids,omitempty" bson:"registration _ids,omitempty"`
	Data         interface{} `json:"data,omitempty" bson:"data,omitempty"`
}
