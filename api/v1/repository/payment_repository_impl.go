package repository

import (
	"shopping-chart/api/v1/helper"
	"shopping-chart/api/v1/model"

	"gorm.io/gorm"
)

type PaymentRepositoryImpl struct {
	db *gorm.DB
}

func NewPayment(db *gorm.DB) PaymentRepository {
	return &PaymentRepositoryImpl{
		db: db,
	}
}

func (pr *PaymentRepositoryImpl) Create(p model.Payment) model.Payment {
	newPayment := model.Payment{
		CartId:     p.CartId,
		CustomerId: p.CustomerId,
		Amount:     p.Amount,
	}
	err := pr.db.Create(&newPayment).Error
	helper.PanicIfError(err)

	return newPayment
}
