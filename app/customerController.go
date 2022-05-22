package app

import (
	"gitlab/awalom/banking/errs"
	"gitlab/awalom/banking/helpers"
	"gitlab/awalom/banking/interfaces"
	"gitlab/awalom/banking/logger"
	"net/http"
)

// CustomerController  Controllers need a service dependency
type CustomerController struct {
	service interfaces.ICustomerService
}

func (s *CustomerController) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	c, err := s.service.GetAllCustomers()
	if err != nil {
		logger.Error("Server error")
		helpers.WriteResponse(w, err.Code, err.AsMsg())
		return
	}
	logger.Info("Fetched customers in the path: " + r.URL.Path)
	helpers.WriteResponse(w, http.StatusOK, c)

}

func (s *CustomerController) GetCustomer(w http.ResponseWriter, r *http.Request) {
	id, _ := helpers.GetVars(r, "id")
	if id == "" {
		logger.Error("id is not provided")
		err := errs.NewUnavailable("id is not provided or it is empty")
		helpers.WriteResponse(w, err.Code, err.AsMsg())
		return
	}
	customer, err := s.service.GetCustomer(id)
	if err != nil {
		logger.Error("Server error")
		helpers.WriteResponse(w, err.Code, err.AsMsg())
		return
	} else {
		logger.Info("Fetched a customer with Id: " + id)
		helpers.WriteResponse(w, http.StatusOK, customer)
	}

}
