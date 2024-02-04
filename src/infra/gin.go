package infra

import (
	"d2-admin-service/src/infra/config"
	"d2-admin-service/src/infra/jwt"
	"d2-admin-service/src/modules/system/dao"
	"d2-admin-service/src/modules/system/rest"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 解决跨域问题
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token,Authorization,Token")
		c.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Content-Type")
		c.Header("Access-Control-Allow-Credentials", "True")
		// 放行options请求
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

//user := new(rest.UserController)
//Routers["/createUser"] = user.CreateUser

func healthCheckApi(c *gin.Context) {
	c.String(http.StatusOK, "online")
}

func RunServer() {
	if config.Config.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	// 默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()
	r.Use(corsMiddleware())
	// 健康检查接口
	r.GET("/", healthCheckApi)
	r.GET("/healthCheck", healthCheckApi)
	// 登入登出接口
	authController := rest.AuthController{}
	r.POST("/login", authController.Login)
	r.POST("/logout", authController.Logout)
	// 定义需要鉴权的接口组
	apiGroup := r.Group("/api/v1")
	apiGroup.Use(jwt.AuthMiddleware())
	{
		routes := dao.ApiDao{}.GetAllRouter()
		for _, route := range routes {
			if route.ApiMethod == "GET" {
				apiGroup.GET(route.ApiPath, InnerRouters[route.ApiName])
				continue
			}
			if route.ApiMethod == "POST" {
				apiGroup.POST(route.ApiPath, InnerRouters[route.ApiName])
			}
		}
	}
	color.Green("============ 启动Rest服务 ============")
	err := r.Run(":" + string(config.Config.Server.Port))
	if err != nil {
		panic("Server StartUp Error: " + err.Error())
	}
}
