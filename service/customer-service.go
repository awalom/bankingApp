package service

import (
	"gitlab/awalom/banking/dta"
	"gitlab/awalom/banking/errs"
	"gitlab/awalom/banking/interfaces"
	"gitlab/awalom/banking/logger"
)

// CustomerService The implementation is a struct
type CustomerService struct {
	Repo interfaces.ICustomerRepository
}

// GetAllCustomers Receiver function
func (s CustomerService) GetAllCustomers() ([]dta.CustomerResponse, *errs.AppError) {

	customers, err := s.Repo.Query()
	var customersResponse []dta.CustomerResponse

	if err != nil {
		logger.Error("could find customers " + err.Message)
		return nil, err
	}
	for _, c := range customers {
		customersResponse = append(customersResponse, dta.ConvertToCustomerResponse(&c))
	}
	return customersResponse, nil
}

func (s CustomerService) GetCustomer(id string) (*dta.CustomerResponse, *errs.AppError) {
	c, err := s.Repo.QueryRow(id)
	if err != nil {
		return &dta.CustomerResponse{}, err

	}
	response := dta.ConvertToCustomerResponse(c)
	return &response, nil
}

// GetCustomerService Helper Function
func GetCustomerService(repository interfaces.ICustomerRepository) CustomerService {
	return CustomerService{repository}
}
