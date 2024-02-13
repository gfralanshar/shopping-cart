package service

import (
	"fmt"
	"shopping-chart/api/v1/helper"
	"shopping-chart/api/v1/model"
	"shopping-chart/api/v1/repository"
	"shopping-chart/api/v1/security"
	"shopping-chart/api/v1/web/dto"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
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

func (cs *CustomerServiceImpl) SignIn(cl dto.LoginCustomerRequestDTO) (dto.LoginCustomerResponseDTO, error) {
	userData := dto.LoginCustomerResponseDTO{}
	customer, err := cs.CustomerRepository.FindUserByUsername(cl.Username)
	helper.PanicIfError(err)

	err = security.VerifyPassword(cl.Password, customer.Password)
	if err == bcrypt.ErrMismatchedHashAndPassword && err != nil {
		return dto.LoginCustomerResponseDTO{}, fmt.Errorf("Unprocessable entity")
	}

	token, err := security.CreateToken(customer.Username)
	helper.PanicIfError(err)

	userData.Token = token
	userData.Username = customer.Username

	return userData, nil
}
