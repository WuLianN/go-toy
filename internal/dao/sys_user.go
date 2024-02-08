package dao

import (
	"time"

	"github.com/WuLianN/go-toy/internal/model"
)

// 用户是否为系统用户
// userName 用户名
// id 用户ID
func (d *Dao) IsSystemUser(userName string, id uint32) (bool, *model.User) {
	user := model.User{UserName: userName, Id: id}
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
func (d *Dao) Register(UserName string, Password string) (uint32, error) {
	createTime := time.Now().Format(time.DateTime)
	user := model.User{UserName: UserName, Password: Password, CreateTime: createTime}
	err := d.engine.Create(&user).Error
	if err != nil {
		return user.Id, err
	}
	return user.Id, nil
}

// 更换密码
func (d *Dao) ChangePassword(userId uint32, passwordHash []byte) bool {
	result := d.engine.Table("user").Where("id = ?", userId).Update("password", passwordHash)

	if result.RowsAffected == 0 || result.Error != nil {
		return false
	}

	return true
}
