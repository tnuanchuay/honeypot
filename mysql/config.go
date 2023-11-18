package mysql

import (
	"github.com/spf13/viper"
	"time"
)

func Config() (string, string, string, int, int, time.Duration) {
	return viper.GetString("db.user"),
		viper.GetString("db.password"),
		viper.GetString("db.dbname"),
		viper.GetInt("db.max_conn"),
		viper.GetInt("db.max_idle_conn"),
		viper.GetDuration("db.conn_max_lifetime")
}

func getTimeout() time.Duration {
	st := viper.GetDuration("db.timeout")
	if st == time.Duration(0) {
		st = 60 * time.Second
	}

	return st
}
