package model

type Account struct {
	AccountId   string  `json:"account-id"`
	CustomerId  string  `json:"customer-id"`
	OpeningDate string  `json:"opening-date"`
	AccountType string  `json:"account-type"`
	Amount      float64 `json:"amount"`
	Status      string  `json:"status"`
}
