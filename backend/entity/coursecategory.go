package entity

import (
	
	"gorm.io/gorm"
	
)

type CourseCategory struct {
	gorm.Model
	CategoryName string

	// 1 Category มีได้หลาย Course
	Courses []Course `gorm:"foreignKey:CourseCategoryID"`
	
}