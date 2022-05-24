package model

type Account struct {
	AccountId   string  `json:"account_id" db:"account_id"`
	CustomerId  string  `json:"customer_id" db:"customer_id"`
	OpeningDate string  `json:"opening_date" db:"opening_date"`
	AccountType string  `json:"account_type" db:"account_type"`
	Amount      float64 `json:"amount" db:"amount"`
	Status      string  `json:"status" db:"status"`
}
