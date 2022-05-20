package repo

import (
	"gitlab/awalom/banking/errs"
	"gitlab/awalom/banking/model"
	"gitlab/awalom/banking/sqlClient"
)

// CustomerRepo  The implementation is a struct
type CustomerRepo struct {
	SqlClient sqlClient.SqlClient
}

// FindAll  Receiver function
func (s CustomerRepo) FindAll() ([]model.Customer, *errs.AppError) {
	return s.SqlClient.Query()
}

func (s CustomerRepo) FindOne(id string) (*model.Customer, *errs.AppError) {
	return s.SqlClient.QueryRow(id)
}

// GetCustomerRepo  Helper Function
func GetCustomerRepo(sqlClient sqlClient.SqlClient) CustomerRepo {
	return CustomerRepo{sqlClient}
}
