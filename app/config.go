package app

import (
	"os"
)

var (
	UserName     string
	Password     string
	DbPort       string
	DBUrl        string
	DatabaseName string
	LogEnv       string
)

const (
	defaultUserName     = "root"
	defaultPassword     = "password"
	defaultDbAdr        = "localhost"
	defaultDbPort       = "3306"
	defaultDatabaseName = "banking"
	defaultAppPort      = ":8081"
	defaultLogEng       = "dev"
)

func init() {

	UserName = os.Getenv("USER_NAME")
	if UserName == "" {
		UserName = defaultUserName
	}
	Password = os.Getenv("PASSWORD")
	if Password == "" {
		Password = defaultPassword
	}
	DBUrl = os.Getenv("URL")
	if DBUrl == "" {
		DBUrl = defaultDbAdr
	}

	DbPort = os.Getenv("DBPORT")
	if DbPort == "" {
		DbPort = defaultDbPort
	}

	DatabaseName = os.Getenv("DATABASE_NAME")
	if DatabaseName == "" {
		DatabaseName = defaultDatabaseName
	}
	LogEnv = os.Getenv("ENV")
	if LogEnv == "" {
		LogEnv = defaultLogEng
	}
}
