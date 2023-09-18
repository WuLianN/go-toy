package model

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

	Title string `json:"title"`
	HideMenu int32 `json:"hide_menu"`
	Icon string `json:"icon"`
	HideChildrenInMenu int32 `json:"hideChildrenInMenu"`

	Group string `json:"group"`
}

func (model Menu) TableName() string {
 return "menu"
}