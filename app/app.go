package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"gitlab/awalom/banking/client"
	"gitlab/awalom/banking/repo"
	"gitlab/awalom/banking/service"
	"log"
	"net/http"
)

func Start() {

	router := mux.NewRouter()
	//wiring
	dbClient := client.GetSqlClient(UserName, Password, DBUrl, DbPort, DatabaseName)
	customerRepo := repo.GetCustomerRepo(dbClient)
	accountRepo := repo.GetAccountRep0(dbClient)
	transactionRepo := repo.GetTransactionRepo(dbClient)
	ch := CustomerController{service.GetCustomerService(customerRepo)}
	ah := AccountController{service.GetAccountService(accountRepo)}
	th := TransactionController{service.GetTransactionService(transactionRepo)}

	router.HandleFunc("/customers", ch.GetAllCustomers)
	router.HandleFunc("/customers/{id:[0-9]+}", ch.GetCustomer)
	router.HandleFunc("/customers/{id}/account", ah.AddNewAccount).Methods(http.MethodPost)
	router.HandleFunc("/customers/account/{account_id}", th.AddNewTransaction).Methods(http.MethodPost)
	fmt.Println("Listening at port 8081 .................")
	log.Fatalln(http.ListenAndServe(defaultAppPort, router))

}
