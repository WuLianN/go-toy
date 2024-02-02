package model

type Menu struct {
	// id
	Id uint32 `json:"id"`
	// 菜单名称
	Name string `json:"name" binding:"required"`
	// 菜单路径
	Path string `json:"path"`
	// component
	Component string `json:"component"`
	// redirect
	Redirect string `json:"redirect"`
	// 父级id
	ParentId uint32 `json:"parent_id"`
	// meta id
	MetaId uint32 `json:"meta_id"`
	// category
	Category string `json:"category"`
	// user id
	UserId uint32 `json:"user_id"`
	// 是否使用
	IsUse uint8 `json:"is_use"`
}

type MenuMeat struct {
	Menu

	Title              string `json:"title"`
	HideMenu           uint32 `json:"hide_menu"`
	Icon               string `json:"icon"`
	HideChildrenInMenu uint32 `json:"hideChildrenInMenu"`
	TagId              uint32 `json:"tag_id"`
}

type Meta struct {
	Id   uint32 `json:"id"`
	Icon string `json:"icon"`
}

func (model Menu) TableName() string {
	return "menu"
}
