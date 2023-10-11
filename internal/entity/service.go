package entity

type Service struct {
	Id         string              `mapstructure:"id"`
	Name       string              `mapstructure:"name"`
	SubService []map[string]string `mapstructure:"sub_services"`
}
