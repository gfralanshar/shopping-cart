package routes

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewRoutes() *httprouter.Router {
	router := httprouter.New()
	router.GET("/hello", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprintf(w, "Hello")
	})

	return router
}
