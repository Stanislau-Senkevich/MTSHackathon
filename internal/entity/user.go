package entity

type User struct {
	Id          string `json:"id" bson:"id"`
	PhoneNumber string `json:"phone_number" bson:"phone_number"`
	Surname     string `json:"surname" bson:"surname"`
}
