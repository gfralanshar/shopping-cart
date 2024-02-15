package service

import "shopping-chart/api/v1/web/dto"

type ProductService interface {
	CreateProduct(productReq dto.CreateProductRequestDTO) dto.CreateProductResponseDTO
}
