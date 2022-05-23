package dto

import "gitlab/awalom/banking/errs"

type TransactionRequest struct {
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
}

func (t TransactionRequest) ValidateTransaction(availableBalance float64) *errs.AppError {

	//if t.TransactionType != "withdrawal" || t.TransactionType != "deposit" {
	//	return errs.ValidationError("Transaction type should be withdrawal or deposit")
	//}
	if t.Amount > availableBalance {
		return errs.ValidationError("Fund it too low.")
	}
	return nil
}
