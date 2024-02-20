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

func ToProductList(products []model.Product) []ProductListResponseDTO {
	productList := []ProductListResponseDTO{}

	for _, product := range products {
		productList = append(productList, ProductListResponseDTO{
			Name:     product.Name,
			Price:    product.Price,
			Category: product.Category.CategoryName,
		})
	}

	return productList
}
