package models

// Car represents teh model for an car
type Car struct {
	Owner   string `gorm:"varchar(100)"`
	Brand   string `gorm:"varchar(100)"`
	Price   int    `gorm:"integer(11)"`
	CarType string `gorm:"varchar(100)"`
	Id      uint   `gorm:"primaryKey"`
}
