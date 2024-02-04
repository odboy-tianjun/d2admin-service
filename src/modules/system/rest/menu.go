package rest

import "github.com/gin-gonic/gin"

type MenuController struct {
}

func (MenuController) TreeMenu(c *gin.Context)   {}
func (MenuController) CreateMenu(c *gin.Context) {}
func (MenuController) DeleteMenu(c *gin.Context) {}
func (MenuController) ModifyMenu(c *gin.Context) {}
