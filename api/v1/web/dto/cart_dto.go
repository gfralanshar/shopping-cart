package dto

type CreateCartRequestDTO struct {
	CustomerId int `json:"customer_id"`
	ProductId  int `json:"product_id"`
	Quantity   int `json:"quantity"`
}

type CreateCartResponseDTO struct {
	CustomerId int `json:"customer_id"`
	ProductId  int `json:"product_id"`
	Quantity   int `json:"quantity"`
}
