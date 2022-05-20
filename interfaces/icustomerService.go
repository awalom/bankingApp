package interfaces

import (
	"gitlab/awalom/banking/errs"
	"gitlab/awalom/banking/model"
)

// This is port. All ports are interfaces

type ICustomerService interface {
	GetAllCustomers() ([]model.Customer, *errs.AppError)
	GetCustomer(id string) (*model.Customer, *errs.AppError)
}
