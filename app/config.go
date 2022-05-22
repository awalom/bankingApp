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
)

const (
	defaultUserName     = "root"
	defaultPassword     = ""
	defaultDbAdr        = "localhost"
	defaultDbPort       = "3306"
	defaultDatabaseName = "banking"
	defaultAppPort      = ":8081"
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
}
