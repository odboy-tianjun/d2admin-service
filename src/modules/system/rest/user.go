package rest

import (
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (UserController) PageUser(c *gin.Context)           {}
func (UserController) CreateUser(c *gin.Context)         {}
func (UserController) DeleteUser(c *gin.Context)         {}
func (UserController) ModifyUserEmail(c *gin.Context)    {}
func (UserController) ModifyUserPassword(c *gin.Context) {}
func (UserController) ModifyUserPhone(c *gin.Context)    {}
