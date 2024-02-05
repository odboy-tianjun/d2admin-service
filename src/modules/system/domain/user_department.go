package domain

import "github.com/jinzhu/gorm"

type UserDepartment struct {
	gorm.Model
	UserId       uint `gorm:"not null"`
	DepartmentId uint `gorm:"not null"`
}

func (UserDepartment) TableName() string {
	return "system_user_department"
}
