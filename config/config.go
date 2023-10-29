package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func load() {
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func Init() {
	load()
}
