package entity

type User struct {
	Id          string `mapstructure:"id"`
	PhoneNumber string `mapstructure:"phone_number"`
	Surname     string `mapstructure:"surname"`
}
