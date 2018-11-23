package model

type Category struct {
	Model
	Name   string `gorm:"size:50"`
	Status int    `gorm:"type:tinyint(1);default:1"`
}
