package entity

import (
	"time"

	"gorm.io/gorm"
)

type Review struct {
    gorm.Model
    Rating       uint
    Comment      string
    ReviewDate   time.Time
    Picture      string `gorm:"type:longtext"`

    UserID       *uint
    User         User `gorm:"foreignKey:UserID"`

    CourseID     *uint
    Course       Course `gorm:"foreignKey:CourseID"`

    Like []Like `gorm:"foreignKey:ReviewID"`
}
