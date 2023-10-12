package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name" gorm:"size:255"`
	Email    string `json:"email" form:"email" gorm:"size:255;unique"`
	Password string `json:"-" form:"password" gorm:"size:255"`
	Token    string `gorm:"-"`
	Tasks    []Task
}
