package service

import (
	"github.com/WuLianN/go-toy/internal/model"
)

type TreeList struct {
	// id
	Id uint32 `json:"id"`
	// 菜单名称
	Name string `json:"name"`
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

	Label string      `json:"label"`
	Tags  []model.Tag `json:"tags"`
}

func (svc *Service) GetMenuList(UserId uint32) []TreeList {
	menus := svc.dao.GetMenu(UserId)

	if menus != nil {
		menuList := []TreeList{}

		for index, menu := range menus {
			tags, _ := svc.dao.QueryMenuTags(menu.Id)
			menus[index].Tags = append(menu.Tags, tags...)
		}

		menuList = append(menuList, GetTreeMenu(menus, 0)...)

		return menuList
	}
	return nil
}

func GetTreeMenu(menuList []model.MenuMeta, pid uint32) []TreeList {
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
				Tags:      v.Tags,
			}
			node.Children = child
			treeList = append(treeList, node)
		}
	}
	return treeList
}

func GetMeta(menu model.MenuMeta) map[string]any {
	meta := make(map[string]any)
	meta["id"] = menu.MetaId

	if menu.Icon != "" {
		meta["icon"] = menu.Icon
	}

	return meta
}

func (svc *Service) AddMenuItem(req model.AddMenuItem, userId uint32) (model.AddMenuItem, error) {
	return svc.dao.AddMenuItem(req.Name, req.ParentId, userId)
}

func (svc *Service) DeleteMenuItem(req model.DeleteMenuItem, userId uint32) error {
	return svc.dao.DeleteMenuItem(req.Id, userId)
}

func (svc *Service) UpdateMenuItem(req *model.UpdateMenuItem) error {
	return svc.dao.UpdateMenuItem(req.Id, req.Name, req.Icon)
}
