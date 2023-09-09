package service

import (
	"golang.org/x/crypto/bcrypt"
)

type UserRequest struct {
	UserName string `form:"user_name" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// 检查登录
func (svc *Service) CheckLogin(param *UserRequest) bool {
	_, hash := svc.dao.IsSystemUser(param.UserName)
	if hash != "" {
		bool := ComparePassword(param.Password, hash)
		return bool
	}

	return false
}

// 检查注册
// @return true 注册成功 false 注册失败
func (svc *Service) CheckRegister(param *UserRequest) (bool, error) {
	isExsited, _ := svc.dao.IsSystemUser(param.UserName)
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