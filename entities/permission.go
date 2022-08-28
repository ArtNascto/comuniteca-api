package entities

import "time"

type Permission struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	DisplayName string

	CreateDate  time.Time
	UpdatedDate time.Time
	DeletedDate *time.Time
}

func (Permission) TableName() string {
	return "public.permission"
}
