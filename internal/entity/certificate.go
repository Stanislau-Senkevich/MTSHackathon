package entity

type Certificate struct {
	Id         string            `mapstructure:"id"`
	OwnerId    string            `mapstructure:"owner_id"`
	Type       string            `mapstructure:"type"`
	Parameters map[string]string `mapstructure:"parameters"`
}
