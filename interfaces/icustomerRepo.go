package interfaces

import (
	"gitlab/awalom/banking/errs"
	"gitlab/awalom/banking/model"
)

type ICustomerRepository interface {
	Query() ([]model.Customer, *errs.AppError)
	QueryRow(id string) (*model.Customer, *errs.AppError)
}
