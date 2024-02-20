package controller

import (
	"net/http"
	"shopping-chart/api/v1/helper"
	"shopping-chart/api/v1/security"
	"shopping-chart/api/v1/service"
	"shopping-chart/api/v1/web"
	"shopping-chart/api/v1/web/dto"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CartControllerImpl struct {
	CartService service.CartService
}

func NewCart(cs service.CartService) CartController {
	return &CartControllerImpl{
		CartService: cs,
	}
}

func (cc *CartControllerImpl) CreateCartHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var req dto.CreateCartRequestDTO

	// take current user id
	id, err := security.ExtractTokenId(r)
	helper.PanicIfError(err)
	req.CustomerId = id

	// take product id
	prodIdParam := p.ByName("product_id")
	prodId, err := strconv.Atoi(prodIdParam)
	helper.PanicIfError(err)
	req.ProductId = prodId

	helper.ReadFromRequestBody(r, &req)

	cart := cc.CartService.AddToCart(req)

	cartResponse := web.WebResponse{
		Status: "created",
		Code:   http.StatusCreated,
		Data:   cart,
	}

	helper.WriteToResponseBody(w, cartResponse)
}
