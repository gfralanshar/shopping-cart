package repository

import (
	"shopping-chart/api/v1/helper"
	"shopping-chart/api/v1/model"

	"gorm.io/gorm"
)

type CartRepositoryImpl struct {
	db *gorm.DB
}

func NewCart(db *gorm.DB) CartRepository {
	return &CartRepositoryImpl{
		db: db,
	}
}

func (cr *CartRepositoryImpl) AddProduct(c model.Cart) model.Cart {
	cart := model.Cart{
		CustomerId: c.CustomerId,
		ProductId:  c.ProductId,
		Quantity:   c.Quantity,
	}
	err := cr.db.Create(&cart).Error
	helper.PanicIfError(err)
	return cart
}
