package model

type Category struct {
	Id           int       `gorm:"column:id;autoIncrement;primary_key"`
	CategoryName string    `gorm:"column:category_name"`
	Products     []Product `gorm:"foreignKey:category_id;references:id"`
}
