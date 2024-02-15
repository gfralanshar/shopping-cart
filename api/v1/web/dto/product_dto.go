package dto

type CreateProductRequestDTO struct {
	Name       string `json:"name" validate:"required"`
	Price      int64  `json:"price"`
	CustomerId int
}

type CreateProductResponseDTO struct {
	Name  string `json:"name" validate:"required"`
	Price int64  `json:"price"`
}
