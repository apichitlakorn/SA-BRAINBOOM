package entity

import (
	"time"
	"gorm.io/gorm"
)

// Task struct to represent a task entity
type Task struct {
	gorm.Model
	Title     string    `gorm:"column:title"`      // Title of the task
	StartDate time.Time `gorm:"column:start_date"` // Start date of the task
	EndDate   time.Time `gorm:"column:end_date"`   // End date of the task
	AllDay    bool      `gorm:"column:all_day"`    // Boolean flag if the task is all-day
	// FK
	UserID *uint `gorm:"column:user_id"`           // Foreign key for user
	User   *User `gorm:"foreignKey:UserID"`        // Reference to User entity
}
