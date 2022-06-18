package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	DbConnStr  string `mapstructure:"connection_string"`
	ServerPort int    `mapstructure:"port"`
}

func LoadAppConfig() (*Config, error) {

	var c *Config

	log.Println("Loading app config")
	viper.AddConfigPath(".") // For now load from root
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
