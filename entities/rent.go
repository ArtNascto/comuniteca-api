package entities

import "time"

type Rent struct {
	ID             uint `gorm:"primaryKey"`
	BookID         uint
	UserId         uint
	Penalty        float64
	RentDate       time.Time
	DevolutionDate *time.Time

	CreateDate  time.Time
	UpdatedDate time.Time
	DeletedDate *time.Time

	Book Book `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	User User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (Rent) TableName() string {
	return "public.rent"
}
