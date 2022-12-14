package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
