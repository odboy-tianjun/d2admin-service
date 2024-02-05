package domain

import "github.com/jinzhu/gorm"

type Api struct {
	gorm.Model
	ApiName   string `gorm:"not null;unique_index:index_npm"` // 接口名称, 例如: getUser
	ApiPath   string `gorm:"not null;unique_index:index_npm"` // 接口路径, 例如: /api/v1/getUser
	ApiMethod string `gorm:"not null;unique_index:index_npm"` // GET、POST
	ApiDesc   string `gorm:"not null"`                        // 接口说明
	ApiStatus uint   `gorm:"not null"`                        // 接口是否可用
}

func (Api) TableName() string {
	return "system_api"
}
