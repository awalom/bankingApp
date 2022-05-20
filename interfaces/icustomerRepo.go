package interfaces

import (
	"gitlab/awalom/banking/errs"
	"gitlab/awalom/banking/model"
)

type ICustomerRepository interface {
	FindAll() ([]model.Customer, *errs.AppError)
	FindOne(id string) (*model.Customer, *errs.AppError)
}
