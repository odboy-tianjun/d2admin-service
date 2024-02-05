package domain

import "github.com/jinzhu/gorm"

type Menu struct {
	gorm.Model
	MenuParentId        uint   `gorm:"null" json:"menuParentId"`
	MenuTitle           string `gorm:"not null" json:"menuTitle"`
	MenuIcon            string `gorm:"not null;default:''" json:"menuIcon"`
	MenuPath            string `gorm:"not null" json:"menuPath"`
	RouterPath          string `gorm:"not null" json:"routerPath"`
	RouterName          string `gorm:"not null" json:"routerName"`
	RouterAuth          uint   `gorm:"not null" json:"routerAuth"`
	RouterHidden        uint   `gorm:"not null" json:"routerHidden"`
	RouterComponentPath string `gorm:"not null" json:"routerComponentPath"`
}

func (Menu) TableName() string {
	return "system_menu"
}
