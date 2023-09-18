package service

import (
	"github.com/WuLianN/go-toy/internal/model"
)

type TreeList struct {
	// id
	Id int32 `json:"id"`
	// 菜单名称
	Name string `json:"name"`
	// 菜单路径
	Path string `json:"path"`
	// component
	Component string `json:"component"`
	// redirect
	Redirect string `json:"redirect"`
	// 父级id
	ParentId int32 `json:"parent_id"`
	// children
	Children []TreeList `json:"children"`
	// meta
	Meta map[string]any `json:"meta"`
}

func (svc *Service) GetMenuList() []TreeList {
	menus := svc.dao.GetMenu()

	if menus != nil {
		// 分组
		systemManagement := make([]model.Menu, 0) // 系统管理
		about := make([]model.Menu, 0) // 关于
		authManagement := make([]model.Menu, 0) // 权限管理

		for _, menu := range menus {
			switch menu.Group {
				case "systemManagement":
					systemManagement = append(systemManagement, menu)
				case "about":
					about = append(about, menu)
				case "authManagement":
					authManagement = append(authManagement, menu)
			}
		}

		systemManagementMenu := GetTreeMenu(systemManagement, 0)
		aboutMenu := GetTreeMenu(about, 0)
		authManagementMenu := GetTreeMenu(authManagement, 0)

		menuList := append(systemManagementMenu, authManagementMenu...)
		menuList = append(menuList, aboutMenu...)

		return menuList
	}
	return nil
}

func GetTreeMenu(menuList []model.Menu, pid int32) []TreeList {
	treeList := []TreeList{}
	for _, v := range menuList {
		if v.ParentId == pid {
			child := GetTreeMenu(menuList, v.Id)
			node := TreeList{
				Id:        v.Id,
				Name:      v.Name,
				Path:      v.Path,
				Component: v.Component,
				Redirect:  v.Redirect,
				ParentId:  v.ParentId,
				Meta:      GetMeta(v),
			}
			node.Children = child
			treeList = append(treeList, node)
		}
	}
	return treeList
}

func GetMeta(menu model.Menu) map[string]any {
	meta := make(map[string]any)
	meta["title"] = menu.Title

	if menu.HideMenu == 1 {
		meta["hideMenu"] = true
	}
	if menu.Icon != "" {
		meta["icon"] = menu.Icon
	}
	if menu.HideChildrenInMenu == 1 {
		meta["hideChildrenInMenu"] = true
	}

	return meta
}