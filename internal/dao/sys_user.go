package dao

import (
	"github.com/WuLianN/go-toy/internal/model"
)

// 用户是否已存在
func (d *Dao) IsSystemUser (UserName string) (bool, *model.User) {
	user := model.User{ UserName: UserName }
	err := d.engine.Table("user").Where("user_name = ?", UserName).First(&user).Error
	if err != nil {
		// fmt.Println(d.engine.Error) // record not found
		d.engine.Error = nil // d.engine.Error设置为nil, 不然下一个sql无法运行, 具体看pkg/opentracing-gorm/otgorm.go, sql追踪造成的
		return false, nil
	}
	return true, &user
}

// 注册
func (d *Dao) Register(UserName string, Password string) (error) {
	user := model.User{ UserName: UserName, Password: Password }
	err := d.engine.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *Dao) GetRoles(userId any) []model.Role {
	var roles []model.Role
	result := d.engine.Table("user_role").Where("user_id = ?", userId).Find(&roles)

	if (result.RowsAffected == 0 || result.Error != nil) {
		return nil
	}

	return roles
}