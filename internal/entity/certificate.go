package entity

type Certificate struct {
	Id         string            `json:"id" bson:"id"`
	OwnerId    string            `json:"owner_id" bson:"owner_id"`
	Type       string            `json:"type" bson:"type"`
	Parameters map[string]string `json:"parameters" bson:"parameters"`
}
