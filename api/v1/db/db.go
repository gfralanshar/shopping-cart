package db

import (
	"fmt"
	"os"
	"shopping-chart/api/v1/helper"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB() *gorm.DB {
	var db *gorm.DB
	var err error

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	// port := os.Getenv("DB_PORT")
	db_name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, db_name)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	// db.AutoMigrate(&model.Customer{}, &model.Product{}, &model.Cart{}, &model.Category{}, &model.Payment{}, &model.CartItems{})

	helper.PanicIfError(err)
	return db
}
