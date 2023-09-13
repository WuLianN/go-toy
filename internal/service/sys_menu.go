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
	Meta map[string]string `json:"meta"`
}

func (svc *Service) GetMenuList() []TreeList {
	menus := svc.dao.GetMenu()

	if menus != nil {
		menuList := GetTreeMenu(menus, 0)

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

func GetMeta(menu model.Menu) map[string]string {
	meta := make(map[string]string)
	meta["title"] = menu.Title

	return meta
}