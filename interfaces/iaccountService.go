package interfaces

import (
	"gitlab/awalom/banking/dta"
	"gitlab/awalom/banking/errs"
)

type IAccountService interface {
	NewAccount(request dta.AccountRequest) (*dta.AccountResponse, *errs.AppError)
}
