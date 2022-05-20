package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"gitlab/awalom/banking/repo"
	"gitlab/awalom/banking/service"
	"gitlab/awalom/banking/sqlClient"
	"log"
	"net/http"
	"os"
)

var (
	UserName     string
	Password     string
	DbProto      string
	DBUrl        string
	DatabaseName string
)

const (
	defaultDatabaseName = "banking"
	defaultUserName     = "root"
	defaultPassword     = ""
	defaultDbPort       = "3306"
	defaultDbAdr        = "localhost"
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

	DbProto = os.Getenv("DBPORT")
	if DbProto == "" {
		DbProto = defaultDbPort
	}

	DBUrl = os.Getenv("URL")
	if DBUrl == "" {
		DBUrl = defaultDbAdr
	}

	DatabaseName = os.Getenv("DATABASE_NAME")
	if DatabaseName == "" {
		DatabaseName = defaultDatabaseName
	}
}

func Start() {

	router := mux.NewRouter()
	//wiring
	ch := CustomerController{service.GetCustomerService(repo.GetCustomerRepo(sqlClient.GetSqlClient(UserName, Password, DbProto, DBUrl, DatabaseName)))}

	router.HandleFunc("/customers", ch.GetAllCustomers)
	router.HandleFunc("/customers/{id:[0-9]+}", ch.GetCustomer)
	fmt.Println("Listening at port 8081 .................")
	log.Fatalln(http.ListenAndServe(defaultAppPort, router))

}
