package entity

type Service struct {
	Id         string              `json:"id" bson:"id"`
	Name       string              `json:"name" bson:"name"`
	SubService []map[string]string `json:"sub_services" bson:"sub_services"`
}
