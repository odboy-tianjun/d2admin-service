package database

import (
	"d2-admin-service/src/infra/config"
	"d2-admin-service/src/modules/system/domain"
	"fmt"
	"github.com/fatih/color"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var DB *gorm.DB

func Connect() {
	color.Green("============ 连接数据库 ============")
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config.Datasource.Username,
		config.Config.Datasource.Password,
		config.Config.Datasource.Host,
		config.Config.Datasource.Port,
		config.Config.Datasource.Database,
	)
	var err error
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic("数据库连接失败, " + err.Error())
	}
	// 连接池
	sqlDB := DB.DB()
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)
	//defer DB.Close()

	// 自动建表
	DB.AutoMigrate(&domain.User{})
	DB.AutoMigrate(&domain.Api{})
	DB.AutoMigrate(&domain.Menu{})
}
