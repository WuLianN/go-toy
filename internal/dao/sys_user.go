package dao

import (
	"github.com/WuLianN/go-toy/internal/model"
)

// 用户是否已存在
func (d *Dao) IsSystemUser (UserName string) (bool, string) {
	user := model.User{ UserName: UserName }
	err := d.engine.Where("user_name = ?", UserName).First(&user).Error
	if err != nil {
		return false, ""
	}
	return true, user.Password
}
