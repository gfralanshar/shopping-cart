package service

import "shopping-chart/api/v1/web/dto"

type PaymentService interface {
	CreatePayment(req dto.CreatePaymentRequestDTO) dto.PaymentResponsesDTO
}
