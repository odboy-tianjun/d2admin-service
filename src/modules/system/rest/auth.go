package rest

import (
	"d2-admin-service/src/infra/context"
	"d2-admin-service/src/infra/jwt"
	"d2-admin-service/src/infra/rediskey"
	"d2-admin-service/src/infra/redistool"
	"d2-admin-service/src/infra/resp"
	"d2-admin-service/src/modules/system/dao"
	dto2 "d2-admin-service/src/modules/system/domain/dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type AuthController struct {
}

// Login 登录
func (AuthController) Login(c *gin.Context) {
	userDao := dao.UserDao{}

	// 声明接收的变量
	var json dto2.LoginDTO
	// 将request的body中的数据，自动按照json格式解析到结构体
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, resp.ResolveParamsError)
		return
	}
	user := userDao.GetUserByUsername(json.Username)
	if user.Username == "" {
		c.JSON(http.StatusBadRequest, resp.UsernameNotExistError)
		return
	}
	// 判断用户名密码是否正确
	if json.Password != user.Password {
		c.JSON(http.StatusBadRequest, resp.PasswordError)
		return
	}
	token := jwt.GenToken(json.Username)
	redistool.Set(rediskey.AUTH_TOKEN+":"+json.Username, token, time.Hour*20)
	context.SetUser(c, user)
	c.JSON(http.StatusOK, resp.Success(map[string]string{
		"username": user.Username,
		"uuid":     user.Uuid,
		"name":     user.Name,
		"token":    token,
	}))
}

// Logout 退出
func (AuthController) Logout(c *gin.Context) {
	user := context.GetUser(c)
	redistool.Remove(rediskey.AUTH_TOKEN + ":" + user.Username)
	c.JSON(http.StatusOK, resp.Success(""))
}

// KickOut 踢出
func (AuthController) KickOut(c *gin.Context) {
	user := context.GetUser(c)
	if user.Username != "admin" {
		c.JSON(http.StatusBadRequest, resp.NoOperationPermissionError)
		return
	}
	var json dto2.KickOutDTO
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, resp.ResolveParamsError)
		return
	}
	redistool.Remove(rediskey.AUTH_TOKEN + ":" + json.Username)
	c.JSON(http.StatusOK, resp.Success(""))
}

func (AuthController) BindUserRole(c *gin.Context) {}
func (AuthController) BindRoleMenu(c *gin.Context) {}
func (AuthController) BindApiMenu(c *gin.Context)  {}
