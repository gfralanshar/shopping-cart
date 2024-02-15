package dto

import (
	"shopping-chart/api/v1/model"
)

func ToProductResponse(p model.Product) CreateProductResponseDTO {
	product := CreateProductResponseDTO{
		Name:  p.Name,
		Price: p.Price,
	}

	return product
}
