package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"gitlab/awalom/banking/repo"
	"gitlab/awalom/banking/service"
	"gitlab/awalom/banking/sqlClient"
	"log"
	"net/http"
)

func Start() {

	router := mux.NewRouter()
	//wiring
	ch := CustomerController{service.GetCustomerService(repo.GetCustomerRepo(sqlClient.GetSqlClient(UserName, Password, DBUrl, DbPort, DatabaseName)))}

	router.HandleFunc("/customers", ch.GetAllCustomers)
	router.HandleFunc("/customers/{id:[0-9]+}", ch.GetCustomer)
	fmt.Println("Listening at port 8081 .................")
	log.Fatalln(http.ListenAndServe(defaultAppPort, router))

}
