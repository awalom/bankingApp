package app

import (
	"gitlab/awalom/banking/errs"
	"gitlab/awalom/banking/helpers"
	"gitlab/awalom/banking/interfaces"
	"net/http"
)

// CustomerController  Controllers need a service dependency
type CustomerController struct {
	service interfaces.ICustomerService
}

func (s *CustomerController) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	c, err := s.service.GetAllCustomers()
	if err != nil {
		helpers.WriteResponse(w, err.Code, err.AsMsg())
		return
	}
	helpers.WriteResponse(w, http.StatusOK, c)

}

func (s *CustomerController) GetCustomer(w http.ResponseWriter, r *http.Request) {
	id, _ := helpers.GetVars(r, "id")
	if id == "" {
		err := errs.NewUnavailable("id is not provided or it is empty")
		helpers.WriteResponse(w, err.Code, err.AsMsg())
		return
	}
	customer, err := s.service.GetCustomer(id)
	if err != nil {
		helpers.WriteResponse(w, err.Code, err.AsMsg())
		return
	} else {
		helpers.WriteResponse(w, http.StatusOK, customer)
	}

}
