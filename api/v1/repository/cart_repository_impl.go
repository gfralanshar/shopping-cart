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

func (cr *CartRepositoryImpl) FindAllCarts(customerId int) []model.Cart {
	carts := []model.Cart{}
	err := cr.db.Model(&model.Cart{}).Preload("Product").Joins("join products on products.id = carts.product_id").Where("carts.customer_id = ?", customerId).Find(&carts).Error
	helper.PanicIfError(err)
	return carts
}
