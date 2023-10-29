package mysql

import "github.com/spf13/viper"

func Config() (string, string, string) {
	return viper.GetString("db.user"), viper.GetString("db.password"), viper.GetString("db.dbname")
}

func getTimeout() string {
	return viper.GetString("db.timeout")
}
