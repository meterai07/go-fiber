package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Fullname string `json:"fullname"`
	Username string `json:"username"`
}
