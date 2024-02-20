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
		Name:       p.Name,
		Price:      p.Price,
		CustomerId: p.CustomerId,
	}

	err := pr.db.Create(&product).Error
	helper.PanicIfError(err)

	return product
}

func (pr *ProductRepositoryImpl) FindProductByCategory(category string) []model.Product {
	var products []model.Product
	var err error
	if category != "" {
		err = pr.db.Model(&model.Product{}).Preload("Category").Joins("join categories on categories.id = products.category_id").Where("categories.category_name = ?", category).Find(&products).Error
	} else {
		err = pr.db.Model(&model.Product{}).Preload("Category").Joins("join categories on categories.id = products.category_id").Find(&products).Error
	}
	helper.PanicIfError(err)
	return products
}

func (pr *ProductRepositoryImpl) FindProductById(id int) model.Product {
	var product model.Product
	err := pr.db.Where("id = ?", id).First(&product).Error
	helper.PanicIfError(err)
	return product
}

func (pr *ProductRepositoryImpl) UpdateProduct(p model.Product) {
	updatedProduct := model.Product{
		Id:       p.Id,
		Quantity: p.Quantity,
	}

	err := pr.db.Model(&model.Product{}).Where("id = ?", p.Id).Update("quantity", updatedProduct.Quantity).Error
	helper.PanicIfError(err)
}
