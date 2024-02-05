package domain

import "github.com/jinzhu/gorm"

type UserRole struct {
	gorm.Model
	UserId uint `gorm:"not null"`
	RoleId uint `gorm:"not null"`
}

func (UserRole) TableName() string {
	return "system_user_role"
}
