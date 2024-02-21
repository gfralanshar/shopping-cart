package model

import (
	"database/sql"
	"time"
)

type Payment struct {
	Id         int          `gorm:"column:id;primary_key;autoIncrement"`
	CartId     int          `gorm:"column:cart_id"`
	CustomerId int          `gorm:"column:customer_id"`
	Amount     int64        `gorm:"column:amount"`
	CreatedAt  time.Time    `gorm:"column:created_at;autoCreateTime;<-create"`
	UpdatedAt  time.Time    `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt  sql.NullTime `gorm:"column:deleted_at"`
}
