package interfaces

import (
	"gitlab/awalom/banking/dta"
	"gitlab/awalom/banking/errs"
)

// This is port. All ports are interfaces

type ICustomerService interface {
	GetAllCustomers() ([]dta.CustomerResponse, *errs.AppError)
	GetCustomer(id string) (*dta.CustomerResponse, *errs.AppError)
}
