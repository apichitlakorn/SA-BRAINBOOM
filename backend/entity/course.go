package entity

import (

	"gorm.io/gorm"
)

type Course struct {
    gorm.Model
    Title  string
    ProfilePicture   string `gorm:"type:longtext"`
    Price  float64
    TeachingPlatform string
	Description string
    Duration uint

    TutorProfileID *uint
	TutorProfile   TutorProfile `gorm:"foreignKey:TutorProfileID"`

    // CourseCategoryID ทำหน้าที่เป็น FK
	CourseCategoryID *uint
	CourseCategory   CourseCategory `gorm:"foreignKey:CourseCategoryID"`

    Reviews []Review `gorm:"foreignKey:CourseID"`
}
