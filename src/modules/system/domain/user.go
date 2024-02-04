package domain

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string
	Password string
}

func (User) TableName() string {
	return "system_user"
}
