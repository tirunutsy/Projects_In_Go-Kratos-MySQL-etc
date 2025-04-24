package models

import (
	"time"
)

type Customer struct {
	ID        uint32    `gorm:"primaryKey"`
	FirstName string    `gorm:"size:100;not null"`
	LastName  string    `gorm:"size:100;not null"`
	Age       int32     `gorm:"type:int"`
	Gender    string    `gorm:"type:string"`
	Emails    []Email   `gorm:"_"`
	Phones    []Phone   `gorm:"_"`
	Addresses []Address `gorm:"_"`
	CreatedAt time.Time `gorm:"index"`
	UpdatedAt time.Time `gorm:"index"`
	// DeletedAt gorm.DeletedAt `gorm:"index"`
}
