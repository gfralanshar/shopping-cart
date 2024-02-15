package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ProductController interface {
	CreateProductHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params )
}
