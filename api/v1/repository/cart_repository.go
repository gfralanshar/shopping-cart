package repository

import "shopping-chart/api/v1/model"

type CartRepository interface {
	Create(cart model.Cart) model.Cart
	AddProduct(cart model.CartItems) model.CartItems
	FindAllCarts(customerId int) []model.CartItems
	FindCartByCustomerId(customerId int) (model.Cart, error)
}
