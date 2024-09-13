package entity

import (
	"time"
	"gorm.io/gorm"
)

type Entity struct {
	gorm.Model
	title     string    
	start_date time.Time 
	end_date   time.Time 
	all_day    bool      
	//FK
	userID *uint
	user   *User `gorm:"foreignKey:UserID"`
}
