package service

import (
	"gitlab/awalom/banking/dto"
	"gitlab/awalom/banking/errs"
	"gitlab/awalom/banking/interfaces"
	"gitlab/awalom/banking/logger"
)

// CustomerService The implementation is a struct
type CustomerService struct {
	Repo interfaces.ICustomerRepository
}

// GetAllCustomers Receiver function
func (s CustomerService) GetAllCustomers() ([]dto.CustomerResponse, *errs.AppError) {

	customers, err := s.Repo.Query()
	var customersResponse []dto.CustomerResponse

	if err != nil {
		logger.Error("could find customers " + err.Message)
		return nil, err
	}
	for _, c := range customers {
		customersResponse = append(customersResponse, dto.ConvertToCustomerResponse(&c))
	}
	return customersResponse, nil
}

func (s CustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.Repo.QueryRow(id)
	if err != nil {
		return &dto.CustomerResponse{}, err

	}
	response := dto.ConvertToCustomerResponse(c)
	return &response, nil
}

// GetCustomerService Helper Function
func GetCustomerService(repository interfaces.ICustomerRepository) CustomerService {
	return CustomerService{repository}
}
