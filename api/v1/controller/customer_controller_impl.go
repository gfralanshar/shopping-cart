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

func (cr *CustomerControllerImpl) LoginHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var loginRequest dto.LoginCustomerRequestDTO
	helper.ReadFromRequestBody(r, &loginRequest)

	userResponse, err := cr.customerService.SignIn(loginRequest)
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Status: "ok",
		Code:   http.StatusOK,
		Data:   userResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}
