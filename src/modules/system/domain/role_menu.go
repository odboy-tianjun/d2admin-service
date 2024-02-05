package domain

import "github.com/jinzhu/gorm"

type RoleMenu struct {
	gorm.Model
	RoleId uint `gorm:"not null"`
	MenuId uint `gorm:"not null"`
}

func (RoleMenu) TableName() string {
	return "system_role_menu"
}
