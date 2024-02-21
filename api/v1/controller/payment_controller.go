package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type PaymentController interface {
	CreatePaymentHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}
