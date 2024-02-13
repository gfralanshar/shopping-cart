package api

import (
	"net/http"
	"os"
	"shopping-chart/api/v1/db"
	"shopping-chart/api/v1/helper"
	"shopping-chart/api/v1/routes"

	"github.com/joho/godotenv"
)

func init() {
	// load a .env file
	err := godotenv.Load()
	helper.PanicIfError(err)

	// connect to db
	db.ConnectDB()
}

func Run() {
	host := os.Getenv("API_HOST")
	port := os.Getenv("API_PORT")
	address := host + ":" + port

	router := routes.NewRoutes()

	server := http.Server{
		Addr:    address,
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
