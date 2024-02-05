package domain

type User struct {
	BaseDomain
	Username string
	Password string
	Uuid     string
	Name     string
}

func (User) TableName() string {
	return "system_user"
}
