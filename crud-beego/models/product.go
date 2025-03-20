package models

import (
	"time"

	"gorm.io/gorm"
)

// "time"

type Product struct {
	ID    int32     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name  string  	`gorm:"column:name;type:varchar(100);not null" json:"name"`
	Price float64 	`gorm:"column:price;type:decimal(10,2);default:0" json:"price"`
	CreatedAt time.Time      `gorm:"column:created_at"` // Akan otomatis terisi saat record dibuat
	UpdatedAt time.Time      `gorm:"column:updated_at"` // Akan otomatis terupdate saat record diubah
	DeletedAt gorm.DeletedAt `gorm:"index"`             // Untuk soft delete
}
