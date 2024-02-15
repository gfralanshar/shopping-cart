package repository

import "shopping-chart/api/v1/model"

type ProductRepository interface {
	Create(product model.Product) model.Product
}
