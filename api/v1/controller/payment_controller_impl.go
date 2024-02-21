package controller

import (
	"net/http"
	"shopping-chart/api/v1/helper"
	"shopping-chart/api/v1/security"
	"shopping-chart/api/v1/service"
	"shopping-chart/api/v1/web"
	"shopping-chart/api/v1/web/dto"

	"github.com/julienschmidt/httprouter"
)

type PaymentControllerImpl struct {
	PaymentService service.PaymentService
}

func NewPayment(paymentService service.PaymentService) PaymentController {
	return &PaymentControllerImpl{
		PaymentService: paymentService,
	}
}

func (pc *PaymentControllerImpl) CreatePaymentHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var req dto.CreatePaymentRequestDTO
	helper.ReadFromRequestBody(r, &req)

	// get current user id
	customerId, err := security.ExtractTokenId(r)
	helper.PanicIfError(err)
	req.CustomerId = customerId

	paymentResponse := pc.PaymentService.CreatePayment(req)
	webResponse := web.WebResponse{
		Status: "created",
		Code:   http.StatusCreated,
		Data:   paymentResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}
