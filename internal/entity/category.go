package entity

type Category struct {
	Id    string `json:"id" bson:"id"`
	Name  string `json:"name" bson:"name"`
	Image string `json:"image" bson:"image"`
	Sum   int    `json:"sum" bson:"sum"`
}
