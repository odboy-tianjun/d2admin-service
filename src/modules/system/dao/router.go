package dao

import (
	"d2-admin-service/src/infra/database"
	"d2-admin-service/src/modules/system/domain"
)

type RouterDao struct {
}

func (RouterDao) GetAllRouter() []domain.Router {
	var routers []domain.Router
	// select * from system_router where router_status = 1
	database.DB.Model(domain.Router{RouterStatus: 1}).Find(&routers)
	return routers
}
