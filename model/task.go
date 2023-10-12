package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	UserID      uint
	Title       string `json:"title" gorm:"size:255"`
	Description string `json:"description" gorm:"type:text"`
	Status      string `json:"status" gorm:"size:50;default:pending"`
}
