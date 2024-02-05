package dao

import (
	"d2-admin-service/src/infra/database"
	"d2-admin-service/src/modules/system/domain"
)

type MenuDao struct {
}

func (MenuDao) GetAllMenu() []domain.Menu {
	var menus []domain.Menu
	// select * from system_menu
	database.DB.Find(&menus)
	return menus
}
