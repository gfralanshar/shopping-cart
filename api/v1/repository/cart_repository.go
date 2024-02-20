package repository

import "shopping-chart/api/v1/model"

type CartRepository interface {
	AddProduct(cart model.Cart) model.Cart
	FindAllCarts(customerId int) []model.Cart
}
