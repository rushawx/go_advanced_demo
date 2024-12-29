package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Phone string `gorm:"uniqueIndex"`
}

type Session struct {
	gorm.Model
	SessionId string
	Code      string
}
