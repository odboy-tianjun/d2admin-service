package dto

type MenuItem struct {
	ID                  uint       `json:"id"`
	MenuParentId        uint       `json:"menuParentId"`
	MenuTitle           string     `json:"menuTitle"`
	MenuIcon            string     `json:"menuIcon"`
	MenuPath            string     `json:"menuPath"`
	RouterPath          string     `json:"routerPath"`
	RouterName          string     `json:"routerName"`
	RouterAuth          uint       `json:"routerAuth"`
	RouterHidden        uint       `json:"routerHidden"`
	RouterComponentPath string     `json:"routerComponentPath"`
	Children            []MenuItem `json:"children"` // 子菜单项
}
