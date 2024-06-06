package config

import (
	"os"
)

type RestConfig struct {
	APIHOST1 string
	APIKEY1  string
}

type DBConfig struct {
	Addr     string
	Port     string
	DB       string
	UserName string
	Password string
}

type Config struct {
	R  RestConfig
	DB DBConfig
}

func Init() Config {
	var r RestConfig
	var d DBConfig
	var c Config
	initStringWithDefVal("APIHOST1", "https://www.omdbapi.com/?i=", &r.APIHOST1)
	initStringWithDefVal("APIKEY1", "e55bdc80", &r.APIKEY1)
	initStringWithDefVal("DBADDR", "172.16.123.174", &d.Addr)
	initStringWithDefVal("DBPORT", "5432", &d.Port)
	initStringWithDefVal("DBNAME", "app_db", &d.DB)
	initStringWithDefVal("DBUSERNAME", "postgres", &d.UserName)
	initStringWithDefVal("DBPASSWORD", "definitelyNotRealPassword", &d.Password)
	c.R = r
	c.DB = d
	return c
}

func initStringWithDefVal(envname string, defval string, variable *string) {
	if _, ok := os.LookupEnv(envname); !ok {
		*variable = defval
		return
	}
	*variable = os.Getenv(envname)
}
