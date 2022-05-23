package interfaces

import (
	"gitlab/awalom/banking/dto"
	"gitlab/awalom/banking/errs"
)

// This is port. All ports are interfaces

type ICustomerService interface {
	GetAllCustomers() ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError)
}
