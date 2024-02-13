package routes

import (
	"shopping-chart/api/v1/controller"

	"github.com/julienschmidt/httprouter"
)

func NewRoutes(customerController controller.CustomerController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/v1/register", customerController.RegisterHandler)
	router.POST("/api/v1/login", customerController.LoginHandler)

	return router
}
