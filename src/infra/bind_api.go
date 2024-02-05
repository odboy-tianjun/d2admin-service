package infra

import (
	"d2-admin-service/src/modules/system/rest"
	"github.com/gin-gonic/gin"
)

// InnerRouters 服务内部路由
var InnerRouters = map[string]gin.HandlerFunc{}

func init() {
	// 用户
	userApi := rest.UserController{}
	InnerRouters["pageUser"] = userApi.PageUser
	InnerRouters["createUser"] = userApi.CreateUser
	InnerRouters["deleteUser"] = userApi.DeleteUser
	InnerRouters["modifyUserEmail"] = userApi.ModifyUserEmail
	InnerRouters["modifyUserPassword"] = userApi.ModifyUserPassword
	InnerRouters["modifyUserPhone"] = userApi.ModifyUserPhone

	// 角色
	roleApi := rest.RoleController{}
	InnerRouters["pageRole"] = roleApi.PageRole
	InnerRouters["createRole"] = roleApi.CreateRole
	InnerRouters["deleteRole"] = roleApi.DeleteRole
	InnerRouters["modifyRoleName"] = roleApi.ModifyRoleName

	// 菜单
	menuApi := rest.MenuController{}
	InnerRouters["queryAllMenus"] = menuApi.TreeMenu
	InnerRouters["createMenu"] = menuApi.CreateMenu
	InnerRouters["deleteMenu"] = menuApi.DeleteMenu
	InnerRouters["modifyMenu"] = menuApi.ModifyMenu

	// api
	api := rest.ApiController{}
	InnerRouters["pageApi"] = api.PageApi
	InnerRouters["registerApi"] = api.RegisterApi
	InnerRouters["unsubscribeApi"] = api.UnSubscribeApi
	InnerRouters["modifyApi"] = api.ModifyApi

	// 关联关系
	authApi := rest.AuthController{}
	InnerRouters["bindUserRole"] = authApi.BindUserRole
	InnerRouters["bindRoleMenu"] = authApi.BindRoleMenu
	InnerRouters["bindApiMenu"] = authApi.BindApiMenu

	// 系统
	InnerRouters["kickOut"] = authApi.KickOut
}
