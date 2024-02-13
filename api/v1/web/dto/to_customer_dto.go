package dto

import "shopping-chart/api/v1/model"

func ToCustomerResponse(c model.Customer) CreateCustomerResponseDTO {
	return CreateCustomerResponseDTO{
		Username: c.Username,
	}
}
