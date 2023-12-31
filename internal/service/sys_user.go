package service

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/WuLianN/go-toy/internal/model"
)

type UserRequest struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserInfoRequest struct {
	UserId uint `json:"userId" binding:"required"`
}

type ChangePasswordRequest struct {
	NewPassword string `json:"newPassword" binding:"required"`
	OldPassword string `json:"oldPassword" binding:"required"`
}

// 检查登录
func (svc *Service) CheckLogin(param *UserRequest) (bool, *model.User) {
	_, userInfo := svc.dao.IsSystemUser(param.UserName, 0)
	if userInfo != nil && userInfo.Password != "" {
		bool := ComparePassword(param.Password, userInfo.Password)
		return bool, userInfo
	}

	return false, nil
}

// 检查注册
// @return true 注册成功 false 注册失败
func (svc *Service) CheckRegister(param *UserRequest) (bool, error) {
	isExsited, _ := svc.dao.IsSystemUser(param.UserName, 0)
	if (isExsited == false) {
		// hash密码
		hash, err := GeneratePassword(param.Password)
		
		if err != nil {
			return false, err // 失败 - 生成hash密码错误
		}

		// 注册写入数据库
		err = svc.dao.Register(param.UserName, string(hash))
		
		if err != nil {
			return false, err // 失败 - 密码入库存储错误
		}

		return true, nil // 成功 - 注册成功
	} else {
		return false, nil // 失败 - 用户已注册
	}
}

// 获取角色权限列表
func (svc *Service) GetRoleList(userId uint) []model.Role{
	list := svc.dao.GetRoles(userId)
	if list != nil {
		return list
	}
	return make([]model.Role, 0)
}

// 生成密码
func GeneratePassword(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return hash, err
}

// 比较密码
func ComparePassword(password string, hash string) (bool) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	} else {
		return true
	}
}

// 更换密码
func (svc *Service) ChangePassword(userId uint, newPassword string) bool {
	newPasswordHash, _ := GeneratePassword(newPassword)
	if len(newPasswordHash) > 0 {
		isSuccessful := svc.dao.ChangePassword(userId, newPasswordHash)
		return isSuccessful
	}
	return false
}