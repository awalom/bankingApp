package interfaces

import (
	"gitlab/awalom/banking/dto"
	"gitlab/awalom/banking/errs"
)

type IAccountService interface {
	NewAccount(request dto.AccountRequest) (*dto.AccountResponse, *errs.AppError)
}
