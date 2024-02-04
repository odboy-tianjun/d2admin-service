package domain

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string
	Password string
	Uuid     string
	Name     string
}

func (User) TableName() string {
	return "system_user"
}
