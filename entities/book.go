package entities

import "time"

type Book struct {
	ID                uint `gorm:"primaryKey"`
	Title             string
	Description       string
	Author            string
	PublishingCompany string
	PublishYear       int
	ISBN              string
	Pages             int
	RentMaxDays       int
	IsAvailable       bool
	CreateDate        time.Time
	UpdatedDate       time.Time
	DeletedDate       *time.Time

	Favorites []FavoriteBook `gorm:"foreignkey:BookID"`
	Comments  []Comment      `gorm:"foreignkey:BookID"`
	Rents     []Rent         `gorm:"foreignkey:BookID"`
}

func (Book) TableName() string {
	return "public.book"
}
