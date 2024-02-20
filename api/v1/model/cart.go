package model

import (
	"database/sql"
	"time"
)

type Cart struct {
	Id         int          `gorm:"column:id;primary_key;autoIncrement"`
	CustomerId int          `gorm:"column:customer_id"`
	ProductId  int          `gorm:"column:product_id"`
	Quantity   int          `gorm:"column:quantity"`
	CreatedAt  time.Time    `gorm:"column:created_at;autoCreateTime;<-create"`
	UpdatedAt  time.Time    `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt  sql.NullTime `gorm:"column:deleted_at"`
	Product    Product      `gorm:"column:product`
}
