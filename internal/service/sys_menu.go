package service

import (
	"github.com/WuLianN/go-toy/internal/model"
)

type TreeList struct {
	// id
	Id uint32 `json:"id"`
	// 菜单名称
	Name  string `json:"name"`
	Label string `json:"label"`
	// 菜单路径
	Path string `json:"path"`
	// component
	Component string `json:"component"`
	// redirect
	Redirect string `json:"redirect"`
	// 父级id
	ParentId uint32 `json:"parent_id"`
	// children
	Children []TreeList `json:"children"`
	// meta
	Meta map[string]any `json:"meta"`
}

func (svc *Service) GetMenuList(UserId uint32) []TreeList {
	menus := svc.dao.GetMenu(UserId)

	if menus != nil {
		// 分类名Map
		categoryNameMap := make(map[string][]model.MenuMeat)

		for _, menu := range menus {
			categoryNameMap[menu.Category] = append(categoryNameMap[menu.Category], menu)
		}

		menuList := []TreeList{}
		for _, category := range categoryNameMap {
			menuList = append(menuList, GetTreeMenu(category, 0)...)
		}

		return menuList
	}
	return nil
}

func GetTreeMenu(menuList []model.MenuMeat, pid uint32) []TreeList {
	treeList := []TreeList{}
	for _, v := range menuList {
		if v.ParentId == pid {
			child := GetTreeMenu(menuList, v.Id)
			node := TreeList{
				Id:        v.Id,
				Name:      v.Name,
				Label:     v.Name,
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

func GetMeta(menu model.MenuMeat) map[string]any {
	meta := make(map[string]any)
	meta["title"] = menu.Title
	meta["tag_id"] = menu.TagId
	meta["id"] = menu.MetaId

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

func (svc *Service) AddMenuItem(req *model.Menu, userId uint32) (uint32, error) {
	return svc.dao.AddMenuItem(req, userId)
}
