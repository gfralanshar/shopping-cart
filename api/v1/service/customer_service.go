package service

import (
	"shopping-chart/api/v1/web/dto"
)

type CustomerService interface {
	Create(customer dto.CreateCustomerRequestDTO) dto.CreateCustomerResponseDTO
	SignIn(customer dto.LoginCustomerRequestDTO) (dto.LoginCustomerResponseDTO, error)
	FindById(customer_id int) dto.CreateCustomerResponseDTO
}
