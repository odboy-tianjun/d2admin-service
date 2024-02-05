package domain

import "github.com/jinzhu/gorm"

type RoleApi struct {
	gorm.Model
	RoleId uint `gorm:"not null"`
	ApiId  uint `gorm:"not null"`
}

func (RoleApi) TableName() string {
	return "system_role_api"
}
