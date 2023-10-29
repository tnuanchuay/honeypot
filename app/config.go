package app

import (
	"fmt"
	"github.com/spf13/viper"
)

func addr() string {
	port := viper.GetString("app.port")
	return fmt.Sprintf(":%s", port)
}
