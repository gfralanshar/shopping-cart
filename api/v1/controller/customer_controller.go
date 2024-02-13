package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CustomerController interface {
	RegisterHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	LoginHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}
