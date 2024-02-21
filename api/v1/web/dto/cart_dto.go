package dto

type CreateCartRequestDTO struct {
	CustomerId int `json:"customer_id"`
}

type CreateCartResponseDTO struct {
	CustomerId int
	CartId   int
}

type AddCartRequestDTO struct {
	CustomerId int
	CartId     int
	ProductId  int
	Quantity   int `json:"quantity"`
}

type AddCartResponseDTO struct {
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`
}

type FindCartByIdDTO struct {
	CustomerId int `json:"customer_id"`
	CartId     int `json:"cart_id"`
}

type ListCartProductDTO struct {
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`
}

type DeleteProductRequestDTO struct {
	CustomerId int
	ProductId int
}
