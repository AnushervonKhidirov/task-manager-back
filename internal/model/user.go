package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"size:30"`
	Surname  string `gorm:"size:30"`
	Email    string
	Password string `json:"-"`
}
