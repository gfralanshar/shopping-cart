package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CartController interface {
	CreateCartHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}
