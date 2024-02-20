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

type ProductControllerImpl struct {
	ProductService  service.ProductService
	CustomerService service.CustomerService
}

func NewProductController(prodService service.ProductService, custService service.CustomerService) ProductController {
	return &ProductControllerImpl{
		ProductService:  prodService,
		CustomerService: custService,
	}
}

func (pc *ProductControllerImpl) CreateProductHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := security.ExtractTokenId(r)
	helper.PanicIfError(err)

	var productReq dto.CreateProductRequestDTO
	helper.ReadFromRequestBody(r, &productReq)
	productReq.CustomerId = id
	productResponse := pc.ProductService.CreateProduct(productReq)
	webResponse := web.WebResponse{
		Status: "created",
		Code:   http.StatusCreated,
		Data:   productResponse,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (pc *ProductControllerImpl) ProductListHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	params := r.URL.Query().Get("category")
	productListResponse := pc.ProductService.ProductListByCategory(params)
	webResponse := web.WebResponse{
		Status: "ok",
		Code:   http.StatusOK,
		Data:   productListResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}
