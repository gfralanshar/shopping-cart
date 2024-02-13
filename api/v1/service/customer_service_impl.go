package service

import (
	"shopping-chart/api/v1/helper"
	"shopping-chart/api/v1/model"
	"shopping-chart/api/v1/repository"
	"shopping-chart/api/v1/security"
	"shopping-chart/api/v1/web/dto"

	"github.com/go-playground/validator/v10"
)

type CustomerServiceImpl struct {
	CustomerRepository repository.CustomerRepository
	Validate           *validator.Validate
}

func NewCustomerService(customerRepository repository.CustomerRepository, validate *validator.Validate) CustomerService {
	return &CustomerServiceImpl{
		CustomerRepository: customerRepository,
		Validate:           validate,
	}
}

func (cs *CustomerServiceImpl) Create(c dto.CreateCustomerRequestDTO) dto.CreateCustomerResponseDTO {
	hashedPassword, err := security.HashPassword(c.Password)
	helper.PanicIfError(err)

	customer := model.Customer{
		Username: c.Username,
		Password: string(hashedPassword),
	}

	newCustomer := cs.CustomerRepository.Create(customer)
	return dto.ToCustomerResponse(newCustomer)
}
