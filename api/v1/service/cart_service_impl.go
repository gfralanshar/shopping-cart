package service

import (
	"shopping-chart/api/v1/exception"
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

	// check if product is present
	product, err := cs.productRepository.FindProductById(req.ProductId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	// check the product owner, make sure if current user not add thier product to cart
	if product.CustomerId == req.CustomerId {
		panic(exception.NewPermissionError("cannot add your own product to the cart"))
	}

	if product.Quantity >= req.Quantity {
		cart := cs.cartRepository.AddProduct(model.CartItems{
			CartId:    req.CartId,
			ProductId: product.Id,
			Quantity:  req.Quantity,
		})

		// update product quantity
		product, err := cs.productRepository.FindProductById(product.Id)
		if err != nil {
			panic(exception.NewNotFoundError(err.Error()))
		}
		updateQuantity := product.Quantity - req.Quantity
		product.Quantity = updateQuantity
		cs.productRepository.UpdateProduct(product)

		return dto.AddCartResponseDTO{
			ProductName: product.Name,
			Quantity:    cart.Quantity,
		}
	} else {
		panic(exception.NewOtherError("cant add product to cart"))
	}
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

func (cs *CartServiceImpl) DeleteProduct(req dto.DeleteProductRequestDTO) {
	// update product quantity
	cart, err := cs.cartRepository.FindCartItemById(req.ProductId, req.CustomerId)
	if err != nil {
		panic(exception.NewNotFoundError("Cart not found"))
	}

	product, err := cs.productRepository.FindProductById(req.ProductId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	cs.cartRepository.DeleteProductByProductId(req.ProductId, req.CustomerId)
	product.Quantity = product.Quantity + cart.Quantity
	cs.productRepository.UpdateProduct(product)
}
