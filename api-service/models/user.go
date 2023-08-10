package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       int    `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Email    string `json:"email" gorm:"unique"`
}
