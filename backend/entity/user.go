package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
    gorm.Model
    Username  string // add
    Password  string
    Email     string
    FirstName string
    LastName  string
    Birthday  time.Time
    Profile   string `gorm:"type:longtext"`

    UserRoleID *uint
	UserRole   UserRole `gorm:"foreignKey:UserRoleID"`

    Reviews []Review `gorm:"foreignKey:UserID"`
}


