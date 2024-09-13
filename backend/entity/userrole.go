package entity

import (
	
	"gorm.io/gorm"
)

type UserRole struct {
	gorm.Model
	RoleName string
	
	// 1 Role มีได้หลาย User
	Users []User `gorm:"foreignKey:UserRoleID"`
}