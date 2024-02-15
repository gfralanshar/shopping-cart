package repository

import (
	"shopping-chart/api/v1/helper"
	"shopping-chart/api/v1/model"

	"gorm.io/gorm"
)

type CustomerRepositoryImpl struct {
	db *gorm.DB
}

func NewCustomer(db *gorm.DB) CustomerRepository {
	return &CustomerRepositoryImpl{
		db: db,
	}
}

func (cr *CustomerRepositoryImpl) Create(c model.Customer) model.Customer {
	customer := model.Customer{
		Username: c.Username,
		Password: c.Password,
	}

	err := cr.db.Create(&customer).Error
	helper.PanicIfError(err)

	return customer
}

func (cr *CustomerRepositoryImpl) FindUserByUsername(username string) (model.Customer, error) {
	customer := model.Customer{}
	err := cr.db.Take(&customer, "username = ?", username).Error
	helper.PanicIfError(err)

	return customer, nil
}

func (cr *CustomerRepositoryImpl) FindUserByCustomerId(id int) (model.Customer, error) {
	customer := model.Customer{}
	err := cr.db.Take(&customer, "id = ?", id).Error
	helper.PanicIfError(err)

	return customer, nil
}
