package config

import (
	"github.com/spf13/viper"
	"os"
)

type db struct {
	Host     string
	User     string
	Password string
	Dbname   string
	Port     string
}

var DB db

func init() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	viper.SetConfigFile(wd + "/.env")
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	DB.Dbname = GetByKey("DB_NAME")
	DB.Host = GetByKey("DB_HOST")
	DB.User = GetByKey("DB_USER")
	DB.Password = GetByKey("DB_PASSWORD")
	DB.Port = GetByKey("DB_PORT")
}

func GetByKey(key string) string {
	return viper.Get(key).(string)
}
