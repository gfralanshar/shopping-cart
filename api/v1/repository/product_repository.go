package repository

import "shopping-chart/api/v1/model"

type ProductRepository interface {
	Create(product model.Product) model.Product
	FindProductByCategory(category string) []model.Product
	FindProductById(id int) (model.Product, error)
	UpdateProduct(product model.Product)
}
