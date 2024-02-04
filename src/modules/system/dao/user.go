package dao

import (
	"d2-admin-service/src/infra/database"
	"d2-admin-service/src/modules/system/domain"
)

type UserDao struct {
}

func (UserDao) GetUserByUsername(username string) *domain.User {
	var user domain.User
	database.DB.Find(&user, "username = ?", username)
	return &user
}
