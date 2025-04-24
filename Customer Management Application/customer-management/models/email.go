package models

import (
	"time"
	//"gorm.io/gorm"
)

type Email struct {
	ID         uint32 `gorm:"primaryKey"`
	Email      string `gorm:"type:varchar(255)"`
	Type       string `gorm:"size:50"`
	IsPrimary  bool
	CustomerID uint32 // Reference to the Customer, but not enforced as a foreign key
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
