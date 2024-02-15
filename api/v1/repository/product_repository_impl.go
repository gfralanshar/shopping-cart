package repository

import (
	"shopping-chart/api/v1/helper"
	"shopping-chart/api/v1/model"

	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{
		db: db,
	}
}

func (pr *ProductRepositoryImpl) Create(p model.Product) model.Product {
	product := model.Product{
		Name:     p.Name,
		Price:    p.Price,
		CustomerId: p.CustomerId,
	}

	err := pr.db.Create(&product).Error
	helper.PanicIfError(err)

	return product
}
