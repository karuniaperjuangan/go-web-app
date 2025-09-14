package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

type RegisterInput struct {
	Username string `json:"username" xml:"username" form:"username"`
	Password string `json:"password" xml:"password" form:"password"`
}

type LoginInput struct {
	Username string `json:"username" xml:"username" form:"username"`
	Password string `json:"password" xml:"password" form:"password"`
}

type UserProfile struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}
