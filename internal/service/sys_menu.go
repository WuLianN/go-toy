package service

import (
	"github.com/WuLianN/go-toy/internal/model"
)

type Menu struct {
	// id 
	Id     int32   `json:"id"`  
	// 菜单名称 
	Name   string  `json:"name"`
	// 菜单路径 
	Path   string  `json:"path"`
	// component
	Component      string  `json:"component"`
	// redirect
	Redirect       string  `json:"redirect"`
	// 父级id
	ParentId       int32   `json:"parent_id"`
	// meta id
	MetaId int32   `json:"meta_id"`

	// children
	Children []model.Menu `json:"children"`
}

func (svc *Service) GetMenuList() []Menu {
	menus := svc.dao.GetMenu()

	parentMenuMap := make(map[int32]Menu) // 父级菜单map

	if (menus != nil) {
		for _, m := range menus {
			if m.ParentId == 0 {
				list := make([]model.Menu, 0)
				parentMenuMap[m.Id] = Menu{ 
					Id: m.Id,
					Name: m.Name,
					Path: m.Path,
					Component: m.Component,
					Redirect: m.Redirect,
					ParentId: m.ParentId,
					MetaId: m.MetaId,
					Children: list, 
				}
			} else {
				parentMenu := parentMenuMap[m.ParentId] // 父级菜单

				parentMenuChildren := parentMenu.Children // 父级菜单的children
				parentMenuChildren = append(parentMenuChildren, m)

				parentMenu.Children = parentMenuChildren
				parentMenuMap[m.ParentId] = parentMenu
			}
		}

		menuList := make([]Menu, 0)

		for _, value := range parentMenuMap {
			menuList = append(menuList, value)
		}
		return menuList
	}
	return nil
}