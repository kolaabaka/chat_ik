package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Sqllite struct {
	URL      string `mapstructure:"url"`
	Password string `mapstructure:"password"`
	Username string `mapstructure:"username"`
	Database string `mapstructure:"database"`
	Type     string `mapstructure:"type"`
}

type DataSource struct {
	Sql Sqllite `mapstructure:"sqllite"`
}

type Config struct {
	DataSource DataSource `mapstructure:"datasource"`
}

func MustInit() Config {
	viper.SetConfigFile("./resources/application.yaml")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Error read file: %v", err))
	}

	var cfg Config

	if err := viper.Unmarshal(&cfg); err != nil {
		panic(fmt.Errorf("Error unmarshaling %v", err))
	}

	return cfg
}
