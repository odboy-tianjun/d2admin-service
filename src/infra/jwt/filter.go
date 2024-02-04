package jwt

import (
	"d2-admin-service/src/infra/context"
	"d2-admin-service/src/infra/resp"
	"d2-admin-service/src/modules/system/domain"
	"d2-admin-service/src/util"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthMiddleware 鉴权中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		color.Green("============ 进入鉴权中间件 ============")
		token := c.Request.Header.Get("Authorization")
		if util.IsBlank(token) {
			c.JSON(http.StatusUnauthorized, resp.NoLoginError)
			c.Abort()
			return
		}
		color.Green("============ 解析token ============")
		if code, claims := ParseToken(token); code == -1 {
			c.JSON(http.StatusUnauthorized, resp.NoLoginError)
			c.Abort()
		} else {
			context.SetUser(c, &domain.User{Username: claims.Username})
			c.Next()
		}
	}
}
