package entities

import "time"

type BookCategory struct {
	ID          uint `gorm:"primaryKey"`
	BookID      uint
	CategoryID  uint
	CreateDate  time.Time
	UpdatedDate time.Time
	DeletedDate *time.Time

	Category Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Book     Book     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (BookCategory) TableName() string {
	return "public.book_category"
}
