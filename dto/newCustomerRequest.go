package dto

import (
	"github.com/andtkach/demogobank/errs"
)

type NewCustomerRequest struct {
	Name        string `json:"name"`
	DateOfBirth string `json:"date_of_birth"`
	City        string `json:"city"`
	ZipCode     string `json:"zip_code"`
}

func (r NewCustomerRequest) Validate() *errs.AppError {
	if r.Name == "" {
		return errs.NewValidationError("Customer name should be provided")
	}
	return nil
}
