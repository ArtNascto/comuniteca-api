package entities

import "time"

type FavoriteBook struct {
	ID          uint `gorm:"primaryKey"`
	BookID      uint
	UserId      uint
	CreateDate  time.Time
	UpdatedDate time.Time
	DeletedDate *time.Time
	Book        Book `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	User        User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (FavoriteBook) TableName() string {
	return "public.favorite_book"
}
