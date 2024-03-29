package dto

type CreateProductRequestDTO struct {
	Name       string `json:"name" validate:"required"`
	Price      int64  `json:"price"`
	CustomerId int
	CategoryId int `json:"category_id" validate:"required"`
}

type CreateProductResponseDTO struct {
	Name  string `json:"name" validate:"required"`
	Price int64  `json:"price"`
}

type ProductListRequestDTO struct {
	Params []string
}

type ProductListResponseDTO struct {
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	Category string `json:"category"`
}
