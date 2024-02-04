package main

import (
	"d2-admin-service/src/infra"
	"d2-admin-service/src/infra/config"
	"d2-admin-service/src/infra/database"
	"d2-admin-service/src/infra/rediscon"
	"d2-admin-service/src/util"
)

func main() {
	config.ImportYmlConfig(util.GetPWD())
	database.Connect()
	rediscon.Connect()
	infra.IntRouters()
	infra.RunServer()
}
