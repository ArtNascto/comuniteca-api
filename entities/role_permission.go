package entities

import "time"

type RolePermission struct {
	ID           uint `gorm:"primaryKey"`
	RoleID       uint
	PermissionID uint

	CreateDate  time.Time
	DeletedDate *time.Time

	Role       Role       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Permission Permission `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (RolePermission) TableName() string {
	return "public.role_permission"
}
