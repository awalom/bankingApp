package app

import (
	"encoding/json"
	"gitlab/awalom/banking/dta"
	"gitlab/awalom/banking/helpers"
	"gitlab/awalom/banking/interfaces"
	"gitlab/awalom/banking/logger"
	"net/http"
)

type AccountController struct {
	accountService interfaces.IAccountService
}

func (c AccountController) AddNewAccount(w http.ResponseWriter, r *http.Request) {
	customerId, err := helpers.GetVars(r, "id")
	if err != nil {
		logger.Warn("Customer Id is not found in the path")
		helpers.WriteResponse(w, http.StatusNotFound, err.Error())
		return
	}

	var requestBody dta.AccountRequest

	err = json.NewDecoder(r.Body).Decode(&requestBody)
	requestBody.CustomerId = customerId
	if err != nil {
		logger.Error("Could not decode request body" + err.Error())
		helpers.WriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	account, appError := c.accountService.NewAccount(requestBody)
	if appError != nil {
		logger.Error("Could not save new Account" + appError.Message)
		helpers.WriteResponse(w, appError.Code, appError.Message)
	} else {
		helpers.WriteResponse(w, http.StatusOK, account)
	}

}
