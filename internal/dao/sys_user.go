package dao

import (
	"errors"
	"time"

	"github.com/WuLianN/go-toy/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
		return false, nil
	}
	return true, &user
}

// 注册
func (d *Dao) Register(UserName string, Password string) (uint32, error) {
	loc, err1 := time.LoadLocation("Asia/Shanghai")

	if err1 != nil {
		loc = time.FixedZone("CST", 8*3600) // 替换上海时间
	}

	createTime := time.Now().In(loc).Format(time.DateTime)
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

func (d *Dao) QueryUser(userId uint32) (model.UserInfo, error) {
	var userInfo model.UserInfo

	if err := d.engine.Table("user").Where("id = ?", userId).Find(&userInfo).Limit(1).Error; err != nil {
		return userInfo, err
	}

	return userInfo, nil
}

func (d *Dao) UpdateUserInfo(userId uint32, userName string, avatar string, isPrivacy uint8) (model.UserInfo, error) {
	userInfo := model.UserInfo{
		UserName:  userName,
		Avatar:    avatar,
		IsPrivacy: isPrivacy,
	}

	if err := d.engine.Table("user").Where("id = ?", userId).Updates(&userInfo).Error; err != nil {
		return userInfo, err
	}

	// updates 不更新零值, isPrivacy有=0的情况, 所以需要单独更新
	if userInfo.IsPrivacy == 0 {
		if err := d.engine.Table("user").Where("id = ?", userId).Updates(map[string]interface{}{"is_privacy": isPrivacy}).Error; err != nil {
			return userInfo, err
		}
	}

	return userInfo, nil
}

func (d *Dao) IsBindedUser(userId1, userId2 uint32) bool {
	binding := model.UserBinding{
		UserId1: userId1,
		UserId2: userId2,
	}
	err := d.engine.Table("user_binding").Where("user_id_1 = ? AND user_id_2 = ?", userId1, userId2).First(&binding).Error

	return !errors.Is(err, gorm.ErrRecordNotFound)
}

func (d *Dao) BindUser(userId1, userId2 uint32) error {
	binding := model.UserBinding{
		UserId1:   userId1,
		UserId2:   userId2,
		CreatedAt: time.Now().Format(time.DateTime),
	}
	err := d.engine.Table("user_binding").Create(&binding).Error

	if err != nil {
		return err
	}
	return nil
}

func (d *Dao) UnbindUser(userId1, userId2 uint32) error {
	binding := model.UserBinding{
		UserId1: userId1,
		UserId2: userId2,
	}
	err := d.engine.Table("user_binding").Where("user_id_1 = ? AND user_id_2 = ?", userId1, userId2).Delete(&binding).Error

	if err != nil {
		return err
	}
	return nil
}

func (d *Dao) QueryBindedUserList(userId uint32) ([]model.UserInfo, error) {
	var list []model.UserInfo
	err := d.engine.Table("user_binding").Select("user.avatar, user.user_name, user.id").Where("user_id_1 = ?", userId).Joins("left join user on user_binding.user_id_2 = user.id").Find(&list).Error

	if err != nil {
		return list, err
	}

	return list, err
}

func (d *Dao) QueryUserSetting(userId uint32) (model.UserSetting, error) {
	userSetting := model.UserSetting{}

	err := d.engine.Table("user_setting").Where("user_id = ?", userId).Find(&userSetting).Limit(1).Error
	if err != nil {
		return userSetting, err
	}

	return userSetting, nil
}

func (d *Dao) UpdateUserSetting(userSetting *model.UserSetting) (model.UserSetting, error) {
	setting := model.UserSetting{
		UserId:       userSetting.UserId,
		PrimaryColor: userSetting.PrimaryColor,
	}

	var err error

	err = d.engine.Table("user_setting").Where("user_id = ?", userSetting.UserId).First(&userSetting).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 没有记录 创建
		err = d.engine.Table("user_setting").Create(&userSetting).Error
	} else {
		// 存在记录 修改
		err = d.engine.Table("user_setting").Clauses(clause.Returning{}).Where("user_id = ?", userSetting.UserId).Updates(&setting).Error
	}

	if err != nil {
		return setting, err
	}
	return setting, nil
}

func (d *Dao) IsAdmin(userId uint32) bool {
	var user model.User
	err := d.engine.Table("user_admin").Where("user_id = ?", userId).Limit(1).Find(&user).Error

	if err != nil {
		return false
	}

	if user.Id > 0 {
		return true
	}

	return false
}
