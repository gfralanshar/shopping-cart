package dto

import (
	"shopping-chart/api/v1/model"
)

func ToCartResponse(c model.Cart) CreateCartResponseDTO {
	return CreateCartResponseDTO{
		CustomerId: c.CustomerId,
		ProductId:  c.ProductId,
		Quantity:   c.Quantity,
	}
}
