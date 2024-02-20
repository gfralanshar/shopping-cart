package service

import "shopping-chart/api/v1/web/dto"

type CartService interface {
	CreateCart(customerId int) dto.CreateCartResponseDTO
	AddToCart(req dto.AddCartRequestDTO) dto.AddCartResponseDTO
	ShowCarts(customerId int) []dto.ListCartProductDTO
	FindCartById(req dto.FindCartByIdDTO) bool
}
