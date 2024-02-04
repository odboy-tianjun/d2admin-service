package config

import (
	"d2-admin-service/src/util"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var Config *ApplicationConfig

type ApplicationConfig struct {
	// 激活哪一个配置文件，格式application-{active}.yml
	Active string
	Debug  bool
	Server struct {
		Port string
	}
	Datasource struct {
		Host     string
		Port     int
		Username string
		Password string
		Database string
	}
	Jwt struct {
		Secret string
	}
	Redis struct {
		Host     string
		Port     int
		Password string
		Database int
		Pool     struct {
			Size int
		}
	}
}

func ImportYmlConfig(currentPath string) {
	parentConfig := readYmlConfigFile(currentPath + string(os.PathSeparator) + "application.yml")
	active := parentConfig.Active
	if util.IsBlank(active) {
		panic("application.yml中的active属性不能为空")
	}
	childrenConfig := readYmlConfigFile(currentPath + string(os.PathSeparator) + "application-" + active + ".yml")
	childrenConfig.Active = parentConfig.Active
	childrenConfig.Debug = parentConfig.Debug
	if util.IsBlank(childrenConfig.Server.Port) {
		// 默认端口
		childrenConfig.Server.Port = "8001"
	}
	Config = childrenConfig
}

// 读取配置文件
func readYmlConfigFile(filepath string) *ApplicationConfig {
	viper.SetConfigFile(filepath)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal errors config file: %w \n", err))
	}
	// 绑定配置到结构体
	var config ApplicationConfig
	viper.Unmarshal(&config)
	return &config
}
