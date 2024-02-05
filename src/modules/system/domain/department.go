package domain

// Department 部门
type Department struct {
	BaseDomain
	DepartmentName     string `gorm:"not null;unique_index:index_npm"` // 部门名称
	DepartmentCode     string `gorm:"not null;unique_index:index_npm"` // 部门编码
	DepartmentDesc     string `gorm:"not null"`
	DepartmentStatus   uint   `gorm:"not null"`
	DepartmentOrder    uint   `gorm:"not null"`
	DepartmentParentId uint   `gorm:"not null"`
}

func (Department) TableName() string {
	return "system_department"
}
