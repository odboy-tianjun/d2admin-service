package domain

import "github.com/jinzhu/gorm"

type Menu struct {
	gorm.Model
	MenuParentId        int    `gorm:"null"`
	MenuTitle           string `gorm:"not null"`
	MenuIcon            string `gorm:"not null;default:''"`
	MenuPath            string `gorm:"not null"`
	RouterPath          string `gorm:"not null"`
	RouterName          string `gorm:"not null"`
	RouterAuth          int    `gorm:"not null"`
	RouterHidden        int    `gorm:"not null"`
	RouterComponentPath string `gorm:"not null"`
}

func (Menu) TableName() string {
	return "system_menu"
}
