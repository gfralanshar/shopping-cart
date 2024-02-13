package service

import (
	"shopping-chart/api/v1/web/dto"
)

type CustomerService interface {
	Create(customer dto.CreateCustomerRequestDTO) dto.CreateCustomerResponseDTO
}
