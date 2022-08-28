package entities

import "time"

type Role struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	DisplayName string

	CreateDate  time.Time
	UpdatedDate time.Time
	DeletedDate *time.Time
}

func (Role) TableName() string {
	return "public.role"
}
