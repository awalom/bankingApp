package interfaces

import (
	"gitlab/awalom/banking/errs"
	"gitlab/awalom/banking/model"
)

type IAccountRepo interface {
	Save(account model.Account) (*model.Account, *errs.AppError)
	GetAccount(id string) (*model.Account, *errs.AppError)
}
