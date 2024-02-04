package dao

import (
	"d2-admin-service/src/infra/database"
	"d2-admin-service/src/modules/system/domain"
)

type RouterDao struct {
}

func (RouterDao) GetAllRouter() []domain.Api {
	var routers []domain.Api
	// select * from system_router where router_status = 1
	database.DB.Model(domain.Api{ApiStatus: 1}).Find(&routers)
	return routers
}
