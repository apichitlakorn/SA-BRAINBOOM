package entity

import (
	

	"gorm.io/gorm"
)

type TutorProfile struct {
	gorm.Model
	Bio  string
	Experience  string
	Education     string
	ProfilePicture string `gorm:"type:longtext"`

	// UserId ทำหน้าที่เป็น FK
	UserID *uint
	User   *User `gorm:"foreignKey:UserID"`
	
}