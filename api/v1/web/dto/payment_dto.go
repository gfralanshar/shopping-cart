package dto

type CreatePaymentRequestDTO struct {
	CustomerId int
	ProductId  []int `json:"product_id"`
}

type DetailProduct struct {
	ProductName string `json:"product_name"`
	Price       int64  `json:"price"`
	Category    string `json:"category"`
	Quantity    int    `json:"quantity"`
}

type PaymentResponsesDTO struct {
	Amount         int64           `json:"amount"`
	ProductDetails []DetailProduct `json:"productDetails"`
}
