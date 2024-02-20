package model

import (
	"database/sql"
	"time"
)

type Customer struct {
	Id        int          `gorm:"column:id;primary_key;autoIncrement"`
	Username  string       `gorm:"column:username;unique"`
	Password  string       `gorm:"column:password"`
	CreatedAt time.Time    `gorm:"column:created_at;autoCreateTime;<-create"`
	UpdatedAt time.Time    `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt sql.NullTime `gorm:"column:deleted_at"`
	Products  []Product    `gorm:"foreignKey:customer_id;references:id"`
	Carts     []Cart       `gorm:"foreignKey:customer_id;references:id"`
}
