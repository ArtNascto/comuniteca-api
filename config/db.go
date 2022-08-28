package config

import (
	"fmt"
	"time"

	"comuniteca.com.br/api/entities"
	"comuniteca.com.br/api/global"
	"comuniteca.com.br/api/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func AssertDatabase() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", global.Config.DB.Host, global.Config.DB.User, global.Config.DB.Pass, global.Config.DB.Name, global.Config.DB.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		utils.Fatal(err)
	}

	global.DB = db
	err = verifyRole()
	if err != nil {
		utils.Fatal(err)
	}
	err = verifyPermissions()
	if err != nil {
		utils.Fatal(err)
	}
	err = verifyAdminRole()
	if err != nil {
		utils.Fatal(err)
	}
	err = verifyAdmin()
	if err != nil {
		utils.Fatal(err)
	}
}

func verifyAdmin() error {

	admin := entities.User{}

	err := global.DB.Where("email = ?", global.Config.Admin.Email).Where("deleted_date is null").First(&admin).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	} else if err != nil && err == gorm.ErrRecordNotFound {
		hashPass, err := global.HashPassword(global.Config.Admin.Password)
		if err != nil {
			return err
		}

		admin = entities.User{
			Name:              global.AdminName,
			Surname:           global.AdminSurname,
			Address:           &global.AdminAddress,
			AddressNumber:     &global.AdminAddressNumber,
			AddressComplement: &global.AdminComplement,
			AddressDistrict:   &global.AdminDistrict,
			AddressCity:       &global.AdminCity,
			AddressState:      &global.AdminState,
			AddressZipcode:    &global.AdminZipcode,
			Email:             global.Config.Admin.Email,
			Password:          hashPass,
			EmailConfirmed:    true,
			CreateDate:        time.Now(),
			UpdatedDate:       time.Now(),
			DeletedDate:       nil,
		}
		err = global.DB.Create(&admin).Error
		if err != nil {
			return err
		}
	}

	return nil
}
func verifyRole() error {
	adminRole := entities.Role{}

	err := global.DB.Where(`name = 'ADMIN' and deleted_date is null`).First(&adminRole).Error
	if err != nil {
		return err
	}

	adminRole = entities.Role{
		Name:        "ADMIN",
		DisplayName: "ADMIN",
		CreateDate:  time.Now(),
		UpdatedDate: time.Now(),
		DeletedDate: nil,
	}
	err = global.DB.Create(&adminRole).Error
	if err != nil {
		return err
	}

	return nil
}

func verifyPermissions() error {
	for _, permission := range global.Permissions {
		entity := entities.Permission{}
		err := global.DB.Where("name = ?", permission).Where("deleted_date is null").First(&entity).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		} else if err != nil && err == gorm.ErrRecordNotFound {
			entity = entities.Permission{
				Name:        permission,
				DisplayName: permission,
				CreateDate:  time.Now(),
				UpdatedDate: time.Now(),
				DeletedDate: nil,
			}
			err = global.DB.Create(&entity).Error
			if err != nil {
				return err
			}

		}
	}

	return nil

}
func verifyAdminRole() error {
	adminRole := entities.Role{}

	err := global.DB.Where(`name = 'ADMIN' and deleted_date is null`).First(&adminRole).Error
	if err != nil {
		return err
	}
	for _, permission := range global.Permissions {
		entity := entities.Permission{}
		err := global.DB.Where("name = ?", permission).Where("deleted_date is null").First(&entity).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}
		rolePermission := entities.RolePermission{}
		err = global.DB.Where("role_id = ?", adminRole.ID).Where("deleted_date is null").First(&rolePermission).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		} else if err != nil && err == gorm.ErrRecordNotFound {
			rolePermission = entities.RolePermission{
				RoleID:       adminRole.ID,
				PermissionID: entity.ID,
				CreateDate:   time.Now(),
				DeletedDate:  nil,
			}
			err = global.DB.Create(&rolePermission).Error
			if err != nil {
				return err
			}
		}
	}

	return nil

}
