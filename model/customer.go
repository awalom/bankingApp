package model

type Customer struct {
	Id          string `db:"customer_id" json:"customer_id"`
	Name        string `db:"name" json:"name"`
	City        string `db:"city" json:"city"`
	Zipcode     string `db:"zipcode" json:"zipcode"`
	DateOfBirth string `db:"date_of_birth" json:"dateOfBirth"`
	Status      string `db:"status" json:"status"`
}

func (c *Customer) StatusAsText() string {
	statusAsString := "active"
	if c.Status == "0" {
		statusAsString = "inactive"
	}
	return statusAsString
}
