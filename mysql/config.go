package mysql

import (
	"github.com/spf13/viper"
	"strconv"
	"time"
)

func Config() (string, string, string) {
	return viper.GetString("db.user"), viper.GetString("db.password"), viper.GetString("db.dbname")
}

func getTimeout() time.Duration {
	st := viper.GetString("db.timeout")
	t, err := strconv.Atoi(st)
	if err != nil {
		return 60 * time.Second
	}

	return time.Duration(t) * time.Millisecond
}
