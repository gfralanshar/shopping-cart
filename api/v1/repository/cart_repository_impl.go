package repository

import (
	"shopping-chart/api/v1/helper"
	"shopping-chart/api/v1/model"
	"time"

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

func (cr *CartRepositoryImpl) Create(c model.Cart) model.Cart {
	cart := model.Cart{
		CustomerId: c.CustomerId,
	}

	err := cr.db.Create(&cart).Error
	helper.PanicIfError(err)

	return cart
}

func (cr *CartRepositoryImpl) AddProduct(c model.CartItems) model.CartItems {
	cartItem := model.CartItems{
		CartId:    c.CartId,
		ProductId: c.ProductId,
		Quantity:  c.Quantity,
	}

	err := cr.db.Preload("Product").Create(&cartItem).Error
	helper.PanicIfError(err)

	return cartItem
}

func (cr *CartRepositoryImpl) FindAllCarts(customerId int) []model.CartItems {
	cartItems := []model.CartItems{}
	err := cr.db.Model(&model.CartItems{}).
		Preload("Product").Joins("join products on products.id = cart_items.product_id").
		Where("customer_id= ?", customerId).
		Find(&cartItems).Error
	helper.PanicIfError(err)
	return cartItems
}

func (cr *CartRepositoryImpl) FindCartByCustomerId(customerId int) (model.Cart, error) {
	var cart model.Cart
	err := cr.db.Where("customer_id", customerId).First(&cart).Error
	if err == gorm.ErrRecordNotFound {
		return model.Cart{}, err
	}

	return cart, nil
}

func (cr *CartRepositoryImpl) FindCartItemById(productId, customerId int) model.CartItems {
	var cartItem model.CartItems
	cr.db.Model(&model.CartItems{}).
		Joins("join carts on carts.id = cart_items.cart_id").
		Where("cart_items.product_id", productId).
		Where("carts.customer_id", customerId).
		Where("cart_items.deleted_at", "NULL").
		First(&cartItem)
	return cartItem
}

func (cr *CartRepositoryImpl) DeleteProductByProductId(productId, customerId int) {
	cartItem := cr.db.Model(&model.CartItems{}).
		Joins("JOIN carts ON carts.id = cart_items.cart_id").
		Where("carts.customer_id = ?", customerId).
		Where("cart_items.product_id = ?", productId).Error
	err := cr.db.Model(&model.CartItems{}).Where("product_id", productId).Delete(&cartItem).Error
	helper.PanicIfError(err)
}

func (cr *CartRepositoryImpl) DeleteCartItemAfterPayment(cartItemId int) {
	cr.db.Model(&model.CartItems{}).Where("id = ?", cartItemId).Update("DeletedAt", time.Now())
}
