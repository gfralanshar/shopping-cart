package dto

import (
	"shopping-chart/api/v1/model"
)

func ToAddToCartResponse(c model.CartItems) AddCartResponseDTO {
	return AddCartResponseDTO{
		ProductName: c.Product.Name,
		Quantity:    c.Quantity,
	}
}

func ToCartsResponses(carts []model.CartItems) []ListCartProductDTO {
	listCartResponses := []ListCartProductDTO{}

	for _, cart := range carts {
		listCartResponse := ListCartProductDTO{
			Quantity:    cart.Quantity,
			ProductName: cart.Product.Name,
		}

		listCartResponses = append(listCartResponses, listCartResponse)
	}

	return listCartResponses
}

func ToCreateCartResponse(cart model.Cart) CreateCartResponseDTO {
	return CreateCartResponseDTO{
		CustomerId: cart.CustomerId,
		CartId: cart.Id,
	}
}
