package controller

import (
	"net/http"
	"shopping-chart/api/v1/helper"
	"shopping-chart/api/v1/service"
	"shopping-chart/api/v1/web"
	"shopping-chart/api/v1/web/dto"

	"github.com/julienschmidt/httprouter"
)

type CustomerControllerImpl struct {
	customerService service.CustomerService
}

func NewCustomerController(customerService service.CustomerService) CustomerController {
	return &CustomerControllerImpl{
		customerService: customerService,
	}
}

func (cc *CustomerControllerImpl) RegisterHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	customerRequest := dto.CreateCustomerRequestDTO{}
	helper.ReadFromRequestBody(r, &customerRequest)

	customerReponse := cc.customerService.Create(customerRequest)
	webResponse := web.WebResponse{
		Status: "OK",
		Code:   http.StatusCreated,
		Data:   customerReponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}
