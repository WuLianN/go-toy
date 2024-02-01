package model

type Menu struct {
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
	// meta id
	MetaId int32 `json:"meta_id"`

	Title              string `json:"title"`
	HideMenu           uint32 `json:"hide_menu"`
	Icon               string `json:"icon"`
	HideChildrenInMenu uint32 `json:"hideChildrenInMenu"`
	TagId              uint32 `json:"tag_id"`

	Category string `json:"category"`
}

func (model Menu) TableName() string {
	return "menu"
}
