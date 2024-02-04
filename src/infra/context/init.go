package context

import (
	"d2-admin-service/src/modules/system/domain"
	"github.com/gin-gonic/gin"
)

func SetUser(c *gin.Context, user *domain.User) {
	c.Set("GenContext_UserInfo", user)
}

func GetUser(c *gin.Context) *domain.User {
	user, _ := c.Get("GenContext_UserInfo")
	return user.(*domain.User)
}
