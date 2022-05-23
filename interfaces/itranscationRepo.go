package interfaces

import (
	"gitlab/awalom/banking/errs"
	"gitlab/awalom/banking/model"
)

type ITransactionRepo interface {
	Save(transaction model.Transaction) (*model.Transaction, *errs.AppError)
}
