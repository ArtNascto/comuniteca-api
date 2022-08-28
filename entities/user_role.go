package entities

import "time"

type UserRole struct {
	ID          uint `gorm:"primaryKey"`
	UserID      uint
	RoleID      uint
	CreateDate  time.Time
	UpdatedDate time.Time
	DeletedDate *time.Time
}

func (UserRole) TableName() string {
	return "public.user_role"
}
