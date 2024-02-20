package model

import (
	"database/sql"
	"time"
)

type CartItems struct {
	Id        int          `gorm:"column:id;primary_key;autoIncrement"`
	CartId    int          `gorm:"column:cart_id"`
	ProductId int          `gorm:"column:product_id"`
	Quantity  int          `gorm:"quantity"`
	Product   Product      `gorm:"foreignKey:product_id;references:id"`
	CreatedAt time.Time    `gorm:"column:created_at;autoCreateTime;<-create"`
	UpdatedAt time.Time    `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt sql.NullTime `gorm:"column:deleted_at"`
}
