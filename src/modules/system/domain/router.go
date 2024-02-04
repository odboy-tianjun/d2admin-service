package domain

import "github.com/jinzhu/gorm"

type Router struct {
	gorm.Model
	RouterName   string `gorm:"not null;unique_index:index_npm"` // 接口名称, 例如: getUser
	RouterPath   string `gorm:"not null;unique_index:index_npm"` // 接口路径, 例如: /api/v1/getUser
	RouterMethod string `gorm:"not null;unique_index:index_npm"` // GET、POST
	RouterDesc   string `gorm:"not null"`                        // 接口说明
	RouterStatus int    `gorm:"not null"`                        // 接口是否可用
}

func (Router) TableName() string {
	return "system_router"
}
