package service

import (
	"gitlab/awalom/banking/errs"
	"gitlab/awalom/banking/interfaces"
	"gitlab/awalom/banking/model"
	"gitlab/awalom/banking/repo"
)

// CustomerService The implementation is a struct
type CustomerService struct {
	Repo interfaces.ICustomerRepository
}

// GetAllCustomers Receiver function
func (s CustomerService) GetAllCustomers() ([]model.Customer, *errs.AppError) {
	return s.Repo.FindAll()
}

func (s CustomerService) GetCustomer(id string) (*model.Customer, *errs.AppError) {
	return s.Repo.FindOne(id)
}

// GetCustomerService Helper Function
func GetCustomerService(repository repo.CustomerRepo) CustomerService {
	return CustomerService{repository}
}
