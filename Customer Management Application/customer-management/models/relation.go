package models

import "time"

type RelationTable struct {
	ID           uint32 `gorm:"primaryKey"`
	CustomerID   uint32
	RelationType string `gorm:"default: 'plan'"`
	RelationID   uint32
	CreatedAt    time.Time
	UpdatedAt    time.Time
	//DeleteddAt   time.Time `gorm:"default:NULL"`
}
