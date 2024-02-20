package model

import (
	"time"
)

type Product struct {
	Id         int       `gorm:"column:id;primary_key;autoIncrement"`
	CategoryId int       `gorm:"column:category_id"`
	Name       string    `gorm:"column:name;size:100"`
	Price      int64     `gorm:"column:price"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime;<-create"`
	UpdatedAt  time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	CustomerId int       `gorm:"column:customer_id"`
	Category   Category  `gorm:"foreignKey:category_id;references:id"`
}
