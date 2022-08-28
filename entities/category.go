package entities

import "time"

type Category struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	DisplayName string
	CreateDate  time.Time
	UpdatedDate time.Time
	DeletedDate *time.Time
}

func (Category) TableName() string {
	return "public.category"
}
