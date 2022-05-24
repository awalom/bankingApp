package app

import (
	"github.com/gorilla/mux"
	"gitlab/awalom/banking/client"
	"gitlab/awalom/banking/logger"
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
	accountRepo := repo.GetAccountRepo(dbClient)
	transactionRepo := repo.GetTransactionRepo(dbClient, accountRepo)
	ch := CustomerController{service.GetCustomerService(customerRepo)}
	ah := AccountController{service.GetAccountService(accountRepo)}
	th := TransactionController{service.GetTransactionService(transactionRepo, accountRepo)}

	router.
		HandleFunc("/customers", ch.GetAllCustomers).
		Methods(http.MethodGet).
		Name("GetAllCustomers")

	router.
		HandleFunc("/customers/{id:[0-9]+}", ch.GetCustomer).
		Methods(http.MethodGet).
		Name("GetCustomer")

	router.
		HandleFunc("/customers/{id}/account", ah.AddNewAccount).
		Methods(http.MethodPost).
		Name("GetAccount")

	router.
		HandleFunc("/customers/account/{account_id}", th.AddNewTransaction).
		Methods(http.MethodPost).
		Name("NewTransaction")

	logger.Info("Listening at port 8081 .................")

	log.Fatalln(http.ListenAndServe(defaultAppPort, router))

}
