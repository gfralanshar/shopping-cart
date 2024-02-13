package routes

import (
	"fmt"
	"net/http"
	"shopping-chart/api/v1/controller"
	"shopping-chart/api/v1/middleware"

	"github.com/julienschmidt/httprouter"
)

func TestMiddleWare(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "Hello World")
}

func NewRoutes(customerController controller.CustomerController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/v1/register", customerController.RegisterHandler)
	router.POST("/api/v1/login", customerController.LoginHandler)
	router.GET("/api/v1/test", middleware.AuthMiddleware(TestMiddleWare))
	return router
}
