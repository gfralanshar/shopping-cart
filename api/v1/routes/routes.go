package routes

import (
	"fmt"
	"net/http"
	"shopping-chart/api/v1/controller"
	"shopping-chart/api/v1/exception"
	"shopping-chart/api/v1/middleware"

	"github.com/julienschmidt/httprouter"
)

func TestMiddleWare(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "Hello World")
}

func NewRoutes(customerController controller.CustomerController, productController controller.ProductController, cartsController controller.CartController, paymentsController controller.PaymentController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/v1/register", customerController.RegisterHandler)
	router.POST("/api/v1/login", customerController.LoginHandler)
	router.GET("/api/v1/test", middleware.AuthMiddleware(TestMiddleWare))

	// products
	router.POST("/api/v1/products", middleware.AuthMiddleware(productController.CreateProductHandler))
	router.GET("/api/v1/products", middleware.AuthMiddleware(productController.ProductListHandler))

	// carts
	router.POST("/api/v1/carts/products/:product_id", middleware.AuthMiddleware(cartsController.CreateCartHandler))
	router.GET("/api/v1/carts", middleware.AuthMiddleware(cartsController.ShowCartsListHandler))
	router.DELETE("/api/v1/carts/products/:product_id", middleware.AuthMiddleware(cartsController.DeleteProductFromCartHandler))

	// payments
	router.POST("/api/v1/payments", middleware.AuthMiddleware(paymentsController.CreatePaymentHandler))

	router.PanicHandler = exception.ErrorHandler

	return router
}
