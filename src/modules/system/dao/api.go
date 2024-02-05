package dao

import (
	"d2-admin-service/src/infra/database"
	"d2-admin-service/src/modules/system/domain"
)

type ApiDao struct {
}

func (ApiDao) GetAllApi() []domain.Api {
	var apis []domain.Api
	// select * from system_api where api_status = 1
	database.DB.Model(domain.Api{ApiStatus: 1}).Find(&apis)
	return apis
}
