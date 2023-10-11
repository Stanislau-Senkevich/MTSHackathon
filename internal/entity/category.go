package entity

type Category struct {
	Id   string `mapstructure:"id"`
	Name string `mapstructure:"name"`
	Sum  int    `mapstructure:"sum"`
}
