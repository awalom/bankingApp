package dto

import (
	"gitlab/awalom/banking/model"
)

type CustomerResponse struct {
	Id          string `json:"customer_id"`
	Name        string `json:"full_name"`
	City        string `json:"city"`
	Zipcode     string `json:"zipcode"`
	DateOfBirth string `json:"date_Of_Birth"`
	Status      string `json:"status"`
	ConvertedBy string `json:"converted_by"`
}

func ConvertToCustomerResponse(customer *model.Customer) CustomerResponse {
	return CustomerResponse{
		Id:          customer.Id,
		Name:        customer.Name,
		City:        customer.City,
		Zipcode:     customer.Zipcode,
		DateOfBirth: customer.DateOfBirth,
		Status:      customer.StatusAsText(),
		ConvertedBy: "Awalom Bereketeab",
	}

}
