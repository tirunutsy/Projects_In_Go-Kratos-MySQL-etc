package models

import "time"

// import (
// 	"time"
// )

type Address struct {
	ID         uint32 `gorm:"primaryKey"`
	Address    string `gorm:"type:varchar(255)"`
	Pincode    string `gorm:"type:varchar(10)"`
	Type       string `gorm:"type:enum('permanant','temperory','other')"` // Match the ENUM definition
	CustomerID uint32
	CreatedAt  time.Time `gorm:"index"`
	UpdatedAt  time.Time `gorm:"index"`
	//DeletedAt  time.Time `gorm:"index"`
}
