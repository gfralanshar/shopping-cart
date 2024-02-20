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

func (cs *CartServiceImpl) CreateCart(customerId int) dto.CreateCartResponseDTO {
	cart := cs.cartRepository.Create(model.Cart{
		CustomerId: customerId,
	})

	return dto.ToCreateCartResponse(cart)
}

func (cs *CartServiceImpl) AddToCart(req dto.AddCartRequestDTO) dto.AddCartResponseDTO {

	product := cs.productRepository.FindProductById(req.ProductId)
	if product.Quantity >= req.Quantity {
		cart := cs.cartRepository.AddProduct(model.CartItems{
			CartId:    req.CartId,
			ProductId: product.Id,
			Quantity:  req.Quantity,
		})

		// update product quantity
		product := cs.productRepository.FindProductById(product.Id)
		updateQuantity := product.Quantity - req.Quantity
		product.Quantity = updateQuantity
		cs.productRepository.UpdateProduct(product)

		return dto.AddCartResponseDTO{
			ProductName: product.Name,
			Quantity:    cart.Quantity,
		}
	}

	return dto.ToAddToCartResponse(model.CartItems{})
}

func (cs *CartServiceImpl) ShowCarts(customerId int) []dto.ListCartProductDTO {
	carts := cs.cartRepository.FindAllCarts(customerId)
	return dto.ToCartsResponses(carts)
}

func (cs *CartServiceImpl) FindCartById(req dto.FindCartByIdDTO) bool {
	_, err := cs.cartRepository.FindCartByCustomerId(req.CustomerId)
	if err == nil {
		return true
	}
	return false
}
