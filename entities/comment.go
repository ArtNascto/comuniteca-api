package entities

import "time"

type Comment struct {
	ID      uint `gorm:"primaryKey"`
	BookID  uint
	UserId  uint
	Title   string
	Comment string
	Stars   float64

	CreateDate  time.Time
	UpdatedDate time.Time
	DeletedDate *time.Time
}

func (Comment) TableName() string {
	return "public.comment"
}
