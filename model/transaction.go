package model

const WITHDRAWAL = "withdrawal"

type Transaction struct {
	TransactionId   string  `db:"transaction_id" db:"transaction_id"`
	AccountId       string  `json:"account_id" db:"account_id"`
	Amount          float64 `json:"amount" db:"amount"`
	TransactionType string  `json:"transaction_type" db:"transaction_type"`
	TransactionDate string  `json:"transaction_date" db:"transaction_date"`
}

func (t Transaction) IsWithdrawal() bool {
	if t.TransactionType == WITHDRAWAL {
		return true
	}
	return false
}
