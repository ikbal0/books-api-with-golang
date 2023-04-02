package models

import "time"

// Car represents teh model for an car
type Book struct {
	ID        uint   `gorm:"primaryKey"`
	NameBook  string `gorm:"varchar(100)"`
	Author    string `gorm:"varchar(100)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
