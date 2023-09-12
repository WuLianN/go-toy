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
	// 1 使用 2不使用
	IsUse  int32   `json:"is_use"`
}

func (model Menu) TableName() string {
 return "menu"
}