package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	DbConnStr  string `mapstructure:"connection_string"`
	ServerPort int    `mapstructure:"port"`
}

func LoadAppConfig() *Config {

	var c *Config

	log.Println("Loading app config")
	viper.AddConfigPath(".") // For now load from root
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Here")
		log.Fatal(err)
	}

	err = viper.Unmarshal(&c)
	if err != nil {
		log.Println("There")
		log.Fatal(err)
	}

	return c
}
