package service

import (
	"github.com/andtkach/demogobank/domain"
	"github.com/andtkach/demogobank/dto"
	"github.com/andtkach/demogobank/errs"
)

type CustomerService interface {
	GetAllCustomers(string) ([]dto.CustomerResponse, *errs.AppError)
	GetOneCustomer(string) (*dto.CustomerResponse, *errs.AppError)
	CreateCustomer(request dto.NewCustomerRequest) (*dto.NewCustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	customers, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}
	response := make([]dto.CustomerResponse, 0)
	for _, c := range customers {
		response = append(response, c.ToDto())
	}
	return response, err
}

func (s DefaultCustomerService) GetOneCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}
	response := c.ToDto()
	return &response, nil
}

func (s DefaultCustomerService) CreateCustomer(req dto.NewCustomerRequest) (*dto.NewCustomerResponse, *errs.AppError) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	customer := domain.NewCustomer(req.Name, req.City, req.ZipCode, req.DateOfBirth)
	if newCustomer, err := s.repo.Save(customer); err != nil {
		return nil, err
	} else {
		return newCustomer.ToNewCustomerResponseDto(), nil
	}
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
