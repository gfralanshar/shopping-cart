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

func (ps *PaymentServiceImpl) CreatePayment(req dto.CreatePaymentRequestDTO) dto.PaymentResponseDTO {
	//select a products cheokout
	var total_amount int64 = 0
	var products []model.Product
	var carts []int

	for _, productId := range req.ProductId {
		// find product
		p, err := ps.productRepository.FindProductById(productId)
		if err != nil {
			panic(exception.NewNotFoundError(err.Error()))
		}
		products = append(products, p)

		// find cart items
		c := ps.cartRepository.FindCartItemById(productId, req.CustomerId)
		amount := p.Price * int64(c.Quantity)
		total_amount += amount
		carts = append(carts, c.CartId)

		ps.cartRepository.DeleteCartItemAfterPayment(c.Id)
		ps.paymentRepository.Create(model.Payment{
			CartId:     c.CartId,
			Amount:     amount,
			CustomerId: req.CustomerId,
		})
	}

	res := dto.PaymentResponseDTO{
		Amount:   total_amount,
		Products: products,
	}

	return res
}
