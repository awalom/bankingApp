package client

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

func GetSqlClient(userName string, password string, dbAddr string, port string, dbname string) *sqlx.DB {
	client, err := sqlx.Open("mysql", getConnectionUrl(userName, password, dbAddr, port, dbname))
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}

func getConnectionUrl(userName string, password string, dbAddr string, port string, dbname string) string {
	dbConnectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", userName, password, dbAddr, port, dbname)
	fmt.Println(dbConnectionString)
	return dbConnectionString
}
