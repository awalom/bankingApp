package app

import (
	"encoding/json"
	"gitlab/awalom/banking/dto"
	"gitlab/awalom/banking/helpers"
	"gitlab/awalom/banking/interfaces"
	"gitlab/awalom/banking/logger"
	"net/http"
)

type TransactionController struct {
	ts interfaces.ITransactionService
}

func (tr TransactionController) AddNewTransaction(w http.ResponseWriter, r *http.Request) {

	//get transaction id
	transactionId, err := helpers.GetVars(r, "account_id")

	if err != nil {
		logger.Error("Error while getting transaction id" + err.Error())
		helpers.WriteResponse(w, http.StatusBadRequest, "Error while getting transaction id")
		return
	}
	//get request body
	var requestBody dto.TransactionRequest
	err = json.NewDecoder(r.Body).Decode(&requestBody)

	if err != nil {
		logger.Error("Error while getting transaction request")
		helpers.WriteResponse(w, http.StatusBadRequest, "Error parsing Decoding request body")
		return
	}
	// add account id to transaction from the path
	requestBody.AccountId = transactionId
	resp, resErr := tr.ts.AddNewTransaction(requestBody)

	if resErr != nil {
		logger.Error("Error while saving new transaction" + resErr.Message)
		helpers.WriteResponse(w, http.StatusInternalServerError, "Error while saving new transaction")
	}

	logger.Info("Transaction was saved" + resp.TransactionId)
	helpers.WriteResponse(w, http.StatusOK, "Transaction was saved"+resp.TransactionId)

}
