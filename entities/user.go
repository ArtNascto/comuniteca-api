package entities

import "time"

type User struct {
	ID                uint `gorm:"primaryKey"`
	Name              string
	Surname           string
	Address           *string
	AddressNumber     *string
	AddressComplement *string
	AddressDistrict   *string
	AddressCity       *string
	AddressState      *string
	AddressZipcode    *string
	Email             string
	Password          string
	EmailConfirmed    bool

	CreateDate  time.Time
	UpdatedDate time.Time
	DeletedDate *time.Time

	FavoriteBooks []FavoriteBook
	Comments      []Comment
	UserRole      UserRole
	Rents         []Rent
}

func (User) TableName() string {
	return "public.user"
}
