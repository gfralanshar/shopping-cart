package service

import (
	"shopping-chart/api/v1/model"
	"shopping-chart/api/v1/repository"
	"shopping-chart/api/v1/web/dto"

	"github.com/go-playground/validator/v10"
)

type CartServiceImpl struct {
	productRepository repository.ProductRepository
	cartRepository    repository.CartRepository
	validator         *validator.Validate
}

func NewCart(pr repository.ProductRepository, cr repository.CartRepository, validate *validator.Validate) CartService {
	return &CartServiceImpl{
		productRepository: pr,
		cartRepository:    cr,
		validator:         validate,
	}
}

func (cs *CartServiceImpl) AddToCart(req dto.CreateCartRequestDTO) dto.CreateCartResponseDTO {

	product := cs.productRepository.FindProductById(req.ProductId)

	if product.Quantity >= req.Quantity {
		cart := cs.cartRepository.AddProduct(model.Cart{
			CustomerId: req.CustomerId,
			ProductId:  req.ProductId,
			Quantity:   req.Quantity,
		})

		// update product quantity
		product := cs.productRepository.FindProductById(req.ProductId)
		updateQuantity := product.Quantity - req.Quantity
		product.Quantity = updateQuantity
		cs.productRepository.UpdateProduct(product)

		return dto.ToCartResponse(cart)
	}

	return dto.ToCartResponse(model.Cart{})
}
