package domain

type Role struct {
	BaseDomain
	RoleName        string
	RoleCode        string
	RoleDescription string
}

func (Role) TableName() string {
	return "system_role"
}
