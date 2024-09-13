package entity

type Like struct {
	UserID   uint  `gorm:"uniqueIndex:user_review_unique"`
	ReviewID *uint `gorm:"uniqueIndex:user_review_unique"`

	Review Review `gorm:"foreignKey:ReviewID"`
}
