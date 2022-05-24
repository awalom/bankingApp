package dto

import (
	"gitlab/awalom/banking/errs"
	"gitlab/awalom/banking/model"
	"time"
)

type AccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (a AccountRequest) ValidateRequest() *errs.AppError {
	if a.Amount < 5000 {
		return errs.ValidationError("To open a new account you need at least 5,000.00")
	}
	if a.AccountType != "saving" && a.AccountType != "checking" {
		return errs.ValidationError("Account type should be checking or saving")
	}
	return nil
}

func (a AccountRequest) AccountRequestToAccount() model.Account {
	return model.Account{
		AccountId:   "",
		CustomerId:  a.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: a.AccountType,
		Amount:      a.Amount,
		Status:      "1",
	}
}
