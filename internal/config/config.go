package config

import (
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
)

const (
	CertificateCollection = "certificate"
	UserCollection        = "user"
	ServiceCollection     = "service"
	CategoryCollection    = "category"
)

type Config struct {
	Port  string
	Mongo MongoConfig
}

type MongoConfig struct {
	User             string
	Password         string
	DBName           string            `mapstructure:"db_name"`
	ConnectionString string            `mapstructure:"conn_string"`
	Collections      map[string]string `mapstructure:"collections"`
}

func InitConfig() (*Config, error) { //nolint
	viper.SetConfigFile("configs/config.yml")
	var cfg Config

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("mongo_config", &cfg.Mongo); err != nil {
		return nil, err
	}

	err := parseEnv(&cfg)

	return &cfg, err
}

func parseEnv(cfg *Config) error {
	if err := gotenv.Load(); err != nil {
		return err
	}

	if err := viper.BindEnv("mongo_user"); err != nil {
		return err
	}

	if err := viper.BindEnv("mongo_password"); err != nil {
		return err
	}

	cfg.Mongo.User = viper.GetString("mongo_user")
	cfg.Mongo.Password = viper.GetString("mongo_password")

	return nil
}
