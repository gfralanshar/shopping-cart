package api

import (
	"net/http"
	"os"
	"shopping-chart/api/v1/controller"
	"shopping-chart/api/v1/db"
	"shopping-chart/api/v1/helper"
	"shopping-chart/api/v1/repository"
	"shopping-chart/api/v1/routes"
	"shopping-chart/api/v1/service"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var dbs *gorm.DB

func init() {
	// load a .env file
	err := godotenv.Load()
	helper.PanicIfError(err)

	// connect to db
	dbs = db.ConnectDB()

	// set log
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
	})
}

func Run() {
	host := os.Getenv("API_HOST")
	port := os.Getenv("API_PORT")
	address := host + ":" + port

	validate := validator.New()
	customerRepository := repository.NewCustomer(dbs)
	customerService := service.NewCustomerService(customerRepository, validate)
	customerController := controller.NewCustomerController(customerService)
	router := routes.NewRoutes(customerController)

	server := http.Server{
		Addr:    address,
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
