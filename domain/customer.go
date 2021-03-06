package domain

import (
	"github.com/andtkach/demogobank/dto"
	"github.com/andtkach/demogobank/errs"
)

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateofBirth string `db:"date_of_birth"`
	Status      string
}

func (c Customer) statusAsText() string {
	statusAsText := "unknown"
	if c.Status == "0" {
		statusAsText = "inactive"
	} else if c.Status == "1" {
		statusAsText = "active"
	}
	return statusAsText
}

func (c Customer) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateofBirth: c.DateofBirth,
		Status:      c.statusAsText(),
	}
}

type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
	Save(customer Customer) (*Customer, *errs.AppError)
}

func NewCustomer(name, city, zipcode, dateOfBirth string) Customer {
	return Customer{
		Name:        name,
		City:        city,
		Zipcode:     zipcode,
		DateofBirth: dateOfBirth,
	}
}

func (c Customer) ToNewCustomerResponseDto() *dto.NewCustomerResponse {
	return &dto.NewCustomerResponse{CustomerId: c.Id}
}
