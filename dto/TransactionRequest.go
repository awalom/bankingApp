package dto

import (
	"gitlab/awalom/banking/errs"
	"gitlab/awalom/banking/model"
	"time"
)

type TransactionRequest struct {
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
}

func (t TransactionRequest) ValidateTransaction() *errs.AppError {

	if t.TransactionType != "withdrawal" && t.TransactionType != "deposit" {
		return errs.ValidationError("Transaction type should be withdrawal or deposit")
	}
	if t.Amount < 10 {
		return errs.ValidationError("Transaction can not be less than 10")
	}
	return nil
}

func (t TransactionRequest) TransactionRequestToTransaction() model.Transaction {
	return model.Transaction{
		TransactionId:   "",
		AccountId:       t.AccountId,
		Amount:          t.Amount,
		TransactionType: t.TransactionType,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
	}
}
