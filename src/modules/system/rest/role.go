package rest

import "github.com/gin-gonic/gin"

type RoleController struct {
}

func (RoleController) PageRole(c *gin.Context)       {}
func (RoleController) CreateRole(c *gin.Context)     {}
func (RoleController) DeleteRole(c *gin.Context)     {}
func (RoleController) ModifyRoleName(c *gin.Context) {}
