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
	c.JSON(http.StatusOK, resp.Success(buildTree(newMenus)))
}
func (MenuController) CreateMenu(c *gin.Context) {}
func (MenuController) DeleteMenu(c *gin.Context) {}
func (MenuController) ModifyMenu(c *gin.Context) {}

// 递归构建无限级树结构
func buildTree(menuItems []dto.MenuItem) []*dto.MenuItem {
	var rootNodes []*dto.MenuItem

	for _, item := range menuItems {
		if item.MenuParentId == 0 { // 根节点
			rootNodes = append(rootNodes, &item)
		} else {
			// 查找并插入到相应父节点的子菜单中
			insertIntoParent(&item, menuItems)
		}
	}

	// 对每个根节点进行递归处理其子节点
	for i := range rootNodes {
		rootNodes[i].Children = buildChildren(rootNodes[i], menuItems)
	}

	return rootNodes
}

// 将一个菜单项插入到其父节点的子菜单列表中（如果父节点已存在于菜单列表中）
func insertIntoParent(item *dto.MenuItem, menuItems []dto.MenuItem) {
	for i := range menuItems {
		if menuItems[i].ID == item.MenuParentId {
			menuItems[i].Children = append(menuItems[i].Children, item)
			return
		}
	}
}

// 递归构建某个节点的所有子节点
func buildChildren(parent *dto.MenuItem, menuItems []dto.MenuItem) []*dto.MenuItem {
	var children []*dto.MenuItem
	processedIds := make(map[uint]bool)

	for _, item := range menuItems {
		if item.MenuParentId == parent.ID && !processedIds[item.ID] { // 防止重复和无限递归
			child := item
			processedIds[child.ID] = true // 标记为已处理

			// 使用剩余未处理的菜单项构建子节点
			child.Children = buildChildren(&child, filterUnprocessed(menuItems, processedIds)) // 递归处理子节点
			children = append(children, &child)
		}
	}

	return children
}

// 创建一个新的切片，其中不包含已处理过的菜单项
func filterUnprocessed(menuItems []dto.MenuItem, processedIds map[uint]bool) []dto.MenuItem {
	var unprocessedItems []dto.MenuItem
	for _, item := range menuItems {
		if !processedIds[item.ID] {
			unprocessedItems = append(unprocessedItems, item)
		}
	}
	return unprocessedItems
}
