package dto

import "shopping-chart/api/v1/model"

type CreatePaymentRequestDTO struct {
	CustomerId int
	ProductId  []int `json:"product_id"`
}

type PaymentResponseDTO struct {
	Amount   int64           `json:"amount"`
	Products []model.Product `json:"products"`
}
