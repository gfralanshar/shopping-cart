package model

import (
	"database/sql"
	"time"
)

type Cart struct {
	Id         int          `gorm:"column:id;primary_key;autoIncrement"`
	CustomerId int          `gorm:"column:customer_id"`
	CreatedAt  time.Time    `gorm:"column:created_at;autoCreateTime;<-create"`
	UpdatedAt  time.Time    `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt  sql.NullTime `gorm:"column:deleted_at"`
	Customer   Customer     `gorm:"foreignKey:customer_id;foreignKey:id"`
	CartItems  []Product    `gorm:"many2many:cart_items;foreignKey:id;joinForeignKey:cart_id;references:id;joinReferences:product_id"`
}
