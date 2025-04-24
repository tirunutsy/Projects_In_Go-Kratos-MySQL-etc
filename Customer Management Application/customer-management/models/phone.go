package models

import (
	"time"
)

// type Phone struct {
// 	ID        uint32    `gorm:"primaryKey"`
// 	Phone     string    `gorm:"size:20;not null;unique"`
// 	IsPrimary bool      `gorm:"default:false"`
// 	Type      string    `gorm:"type:ENUM('work', 'personal', 'other');not null"`
// 	CreatedAt time.Time `gorm:"index"`
// 	UpdatedAt time.Time `gorm:"index"`
// 	DeletedAt time.Time `gorm:"index"`
// }

type Phone struct {
	ID         uint32 `gorm:"primaryKey"`
	Phone      string `gorm:"type:varchar(255)"`
	Type       string
	IsPrimary  bool
	CustomerID uint32 // Reference to the Customer, but not enforced as a foreign key
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
