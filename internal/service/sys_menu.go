package service

import (
	"strconv"

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

	Label string `json:"label"`
}

func (svc *Service) GetMenuList(UserId uint32) []TreeList {
	menus := svc.dao.GetMenu(UserId)

	if menus != nil {
		// 分类名Map
		categoryNameMap := make(map[string][]model.MenuMeat)

		for _, menu := range menus {
			categoryIdStr := strconv.Itoa(int(menu.CategoryId))
			categoryNameMap[categoryIdStr] = append(categoryNameMap[categoryIdStr], menu)
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
	meta["id"] = menu.MetaId
	meta["category_id"] = menu.CategoryId

	if menu.Icon != "" {
		meta["icon"] = menu.Icon
	}

	return meta
}

func (svc *Service) AddMenuItem(req model.AddMenuItem, userId uint32) (model.AddMenuItem, error) {
	return svc.dao.AddMenuItem(req.Name, req.ParentId, req.CategoryId, userId)
}
