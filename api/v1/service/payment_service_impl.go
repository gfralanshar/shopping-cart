package service

import (
	"shopping-chart/api/v1/exception"
	"shopping-chart/api/v1/model"
	"shopping-chart/api/v1/repository"
	"shopping-chart/api/v1/web/dto"

	"github.com/go-playground/validator/v10"
)

type PaymentServiceImpl struct {
	paymentRepository repository.PaymentRepository
	productRepository repository.ProductRepository
	cartRepository    repository.CartRepository
	Validate          *validator.Validate
}

func NewPayment(paymentRepo repository.PaymentRepository, productRepo repository.ProductRepository, cartRepo repository.CartRepository, validate *validator.Validate) PaymentService {
	return &PaymentServiceImpl{
		paymentRepository: paymentRepo,
		productRepository: productRepo,
		cartRepository:    cartRepo,
		Validate:          validate,
	}
}

func (ps *PaymentServiceImpl) CreatePayment(req dto.CreatePaymentRequestDTO) dto.PaymentResponsesDTO {
	//select a products cheokout
	var total_amount int64 = 0
	var products []dto.DetailProduct
	var carts []int

	for _, productId := range req.ProductId {
		// find product
		p, err := ps.productRepository.FindProductById(productId)
		if err != nil {
			panic(exception.NewNotFoundError(err.Error()))
		}

		// find cart items
		c, err := ps.cartRepository.FindCartItemById(productId, req.CustomerId)
		if err != nil {
			exception.NewNotFoundError("cart not found")
		}

		// calculate amount
		amount := p.Price * int64(c.Quantity)
		total_amount += amount
		carts = append(carts, c.CartId)

		// add details of product
		detailProduct := dto.DetailProduct{
			ProductName: p.Name,
			Category:    p.Category.CategoryName,
			Price:       p.Price,
			Quantity:    c.Quantity,
		}
		products = append(products, detailProduct)

		ps.cartRepository.DeleteCartItemAfterPayment(c.Id)
		ps.paymentRepository.Create(model.Payment{
			CartId:     c.CartId,
			Amount:     amount,
			CustomerId: req.CustomerId,
		})
	}

	res := dto.PaymentResponsesDTO{
		Amount:         total_amount,
		ProductDetails: products,
	}

	return res
}
