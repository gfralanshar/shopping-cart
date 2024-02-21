package service

import (
	"shopping-chart/api/v1/helper"
	"shopping-chart/api/v1/model"
	"shopping-chart/api/v1/repository"
	"shopping-chart/api/v1/web/dto"

	"github.com/go-playground/validator/v10"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	Validate          *validator.Validate
}

func NewProductService(prodRepository repository.ProductRepository, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepository: prodRepository,
		Validate:          validate,
	}
}

func (ps *ProductServiceImpl) CreateProduct(p dto.CreateProductRequestDTO) dto.CreateProductResponseDTO {
	err := ps.Validate.Struct(&p)
	helper.PanicIfError(err)

	product := model.Product{
		Name:       p.Name,
		Price:      p.Price,
		CustomerId: p.CustomerId,
		CategoryId: p.CategoryId,
	}

	p2 := ps.ProductRepository.Create(product)
	return dto.ToProductResponse(p2)
}

func (ps *ProductServiceImpl) ProductListByCategory(category string) []dto.ProductListResponseDTO {
	products := ps.ProductRepository.FindProductByCategory(category)
	return dto.ToProductList(products)
}
