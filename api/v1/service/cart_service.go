package service

import "shopping-chart/api/v1/web/dto"

type CartService interface {
	AddToCart(req dto.CreateCartRequestDTO) dto.CreateCartResponseDTO
}
