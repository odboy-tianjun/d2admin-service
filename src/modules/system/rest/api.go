package rest

import "github.com/gin-gonic/gin"

type ApiController struct {
}

func (ApiController) PageApi(c *gin.Context)        {}
func (ApiController) RegisterApi(c *gin.Context)    {}
func (ApiController) UnSubscribeApi(c *gin.Context) {}
func (ApiController) ModifyApi(c *gin.Context)      {}
