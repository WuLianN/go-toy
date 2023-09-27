package dao

import (
	"github.com/WuLianN/go-toy/internal/model"
	"time"
)

// 用户是否为系统用户
// userName 用户名
// id 用户ID
func (d *Dao) IsSystemUser (userName string, id uint) (bool, *model.User) {
	user := model.User{ UserName: userName, Id: id }
	var err error
	if userName != "" {
		err = d.engine.Table("user").Where("user_name = ?", userName).First(&user).Error
	} else if id != 0 {
		err = d.engine.Table("user").Where("id = ?", id).First(&user).Error
	}
	
	if err != nil {
		// fmt.Println(d.engine.Error) // record not found
		d.engine.Error = nil // d.engine.Error设置为nil, 不然下一个sql无法运行, 具体看pkg/opentracing-gorm/otgorm.go, sql追踪造成的
		return false, nil
	}
	return true, &user
}

// 注册
func (d *Dao) Register(UserName string, Password string) (error) {
	createTime := time.Now().Format(time.DateTime)
	user := model.User{ UserName: UserName, Password: Password, CreateTime: createTime }
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

// 更换密码
func (d *Dao) ChangePassword(userId uint, passwordHash []byte) bool {
	result := d.engine.Table("user").Where("id = ?", userId).Update("password", passwordHash)

	if (result.RowsAffected == 0 || result.Error != nil) {
		return false
	}

	return true
}