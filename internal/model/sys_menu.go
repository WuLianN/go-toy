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
	// user id
	UserId uint32 `json:"user_id"`
	// 是否使用
	IsUse uint8 `json:"is_use"`
	// 排序
	Sort uint32 `json:"sort"`
	// 是否私密
	IsPrivacy uint8 `json:"is_privacy"`
}

type MenuMeta struct {
	Menu
	Tags []Tag  `json:"tags" gorm:"foreignKey:Id"`
	Icon string `json:"icon"`
}

type Meta struct {
	Id   uint32 `json:"id"`
	Icon string `json:"icon"`
}

type AddMenuItem struct {
	Id       uint32 `json:"id"`
	Name     string `json:"name"`
	ParentId uint32 `json:"parent_id"`
}

type DeleteMenuItem struct {
	Id uint32 `json:"id" binding:"required"`
}

type UpdateMenuItem struct {
	Id        uint32 `json:"id" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Icon      string `json:"icon"`
	IsUse     uint8  `json:"is_use"`
	IsPrivacy uint8  `json:"is_privacy"`
}

type SaveMenuSort struct {
	ParentId uint32 `json:"parent_id"`
	Id       uint32 `json:"id" binding:"required"`
	Sort     uint32 `json:"sort" binding:"required"`
}

type MenuTags struct {
	Tags   []Tag  `json:"tags" binding:"required"`
	MenuId uint32 `json:"menu_id" binding:"required"`
}

type MenuTag struct {
	TagId  uint32 `json:"tag_id"`
	MenuId uint32 `json:"menu_id"`
}

func (model Menu) TableName() string {
	return "menu"
}
