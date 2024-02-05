package rest

import (
	"d2-admin-service/src/infra/resp"
	"d2-admin-service/src/modules/system/dao"
	"d2-admin-service/src/modules/system/domain/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MenuController struct {
}

func (MenuController) TreeMenu(c *gin.Context) {
	menuDao := dao.MenuDao{}
	menus := menuDao.GetAllMenu()

	var newMenus []dto.MenuItem
	for _, menu := range menus {
		menuItem := dto.MenuItem{}
		menuItem.ID = menu.ID
		menuItem.MenuParentId = menu.MenuParentId
		menuItem.MenuTitle = menu.MenuTitle
		menuItem.MenuIcon = menu.MenuIcon
		menuItem.MenuPath = menu.MenuPath
		menuItem.RouterAuth = menu.RouterAuth
		menuItem.RouterHidden = menu.RouterHidden
		menuItem.RouterName = menu.RouterName
		menuItem.RouterPath = menu.RouterPath
		menuItem.RouterComponentPath = menu.RouterComponentPath
		newMenus = append(newMenus, menuItem)
	}
	c.JSON(http.StatusOK, resp.Success(buildTree(newMenus, 0)))
}
func (MenuController) CreateMenu(c *gin.Context) {}
func (MenuController) DeleteMenu(c *gin.Context) {}
func (MenuController) ModifyMenu(c *gin.Context) {}

// buildTree 获取菜单
func buildTree(menuList []dto.MenuItem, pid uint) []dto.MenuItem {
	var treeList []dto.MenuItem
	for _, v := range menuList {
		if v.MenuParentId == pid {
			child := buildTree(menuList, v.ID)
			node := dto.MenuItem{
				ID:                  v.ID,
				MenuTitle:           v.MenuTitle,
				RouterName:          v.RouterName,
				RouterPath:          v.RouterPath,
				RouterAuth:          v.RouterAuth,
				RouterHidden:        v.RouterHidden,
				RouterComponentPath: v.RouterComponentPath,
				MenuPath:            v.MenuPath,
				MenuIcon:            v.MenuIcon,
				MenuParentId:        v.MenuParentId,
			}
			if len(child) > 0 {
				node.Children = child
			}
			node.Children = child
			treeList = append(treeList, node)
		}
	}
	return treeList
}
