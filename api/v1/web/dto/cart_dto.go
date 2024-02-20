package dto

import "shopping-chart/api/v1/model"

type CreateCartRequestDTO struct {
	CustomerId int `json:"customer_id"`
	ProductId  int `json:"product_id"`
	Quantity   int `json:"quantity"`
}

type CreateCartResponseDTO struct {
	CustomerId int             `json:"customer_id"`
	ProductId  int             `json:"product_id"`
	Quantity   int             `json:"quantity"`
	Products   []model.Product `json:"products`
}

type ListCartProductsDTO struct {
	ProductName string `json:"product_name"`
	Quantity    int `json:"quantity"`
}
