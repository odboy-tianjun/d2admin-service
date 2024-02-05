package domain

import (
	"github.com/jinzhu/gorm"
)

type BaseDomain struct {
	gorm.Model
	CreatedBy string `gorm:"column:created_by;not null"`
	UpdatedBy string `gorm:"column:updated_by;not null"`
}
