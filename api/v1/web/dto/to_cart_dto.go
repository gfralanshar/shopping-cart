package dto

import (
	"shopping-chart/api/v1/model"
)

func ToCartResponse(c model.Cart) CreateCartResponseDTO {
	return CreateCartResponseDTO{
		CustomerId: c.CustomerId,
		ProductId:  c.ProductId,
		Quantity:   c.Quantity,
	}
}

func ToCartsResponses(carts []model.Cart) []ListCartProductsDTO {
	listCartResponses := []ListCartProductsDTO{}

	for _, cart := range carts {
		listCartResponse := ListCartProductsDTO{
			ProductName: cart.Product.Name,
			Quantity:    cart.Quantity,
		}

		listCartResponses = append(listCartResponses, listCartResponse)
	}

	return listCartResponses
}
