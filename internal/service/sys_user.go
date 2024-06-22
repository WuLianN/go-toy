package service

import (
	"errors"

	"github.com/WuLianN/go-toy/internal/model"
	"golang.org/x/crypto/bcrypt"
)

type UserRequest struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserInfoRequest struct {
	Id        uint32 `json:"id"`
	UserName  string `json:"user_name"`
	Avatar    string `json:"avatar"`
	IsPrivacy uint8  `json:"is_privacy"`
}

type ChangePasswordRequest struct {
	NewPassword string `json:"newPassword" binding:"required"`
	OldPassword string `json:"oldPassword" binding:"required"`
}

type UserIdRequest struct {
	Id uint32 `json:"id" binding:"required"`
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
func (svc *Service) CheckRegister(param *UserRequest) (bool, uint32, error) {
	isExsited, _ := svc.dao.IsSystemUser(param.UserName, 0)
	if !isExsited {
		// hash密码
		hash, err := GeneratePassword(param.Password)

		if err != nil {
			return false, 0, err // 失败 - 生成hash密码错误
		}

		// 注册写入数据库
		userId, err2 := svc.dao.Register(param.UserName, string(hash))

		if err2 != nil {
			return false, 0, err2 // 失败 - 密码入库存储错误
		}

		return true, userId, nil // 成功 - 注册成功
	}

	return false, 0, nil // 失败 - 用户已注册
}

// 生成密码
func GeneratePassword(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return hash, err
}

// 比较密码
func ComparePassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	} else {
		return true
	}
}

// 更换密码
func (svc *Service) ChangePassword(userId uint32, newPassword string) bool {
	newPasswordHash, _ := GeneratePassword(newPassword)
	if len(newPasswordHash) > 0 {
		isSuccessful := svc.dao.ChangePassword(userId, newPasswordHash)
		return isSuccessful
	}
	return false
}

func (svc *Service) GetUserInfo(userId uint32) (model.UserInfo, error) {
	return svc.dao.QueryUser(userId)
}

func (svc *Service) UpdateUserInfo(req *UserInfoRequest) (model.UserInfo, error) {
	return svc.dao.UpdateUserInfo(req.Id, req.UserName, req.Avatar, req.IsPrivacy)
}

func (svc *Service) BindUser(userId uint32, req *UserRequest) error {
	var err error

	isExsited, userInfo := svc.dao.IsSystemUser(req.UserName, 0)

	if !isExsited {
		return errors.New("用户不存在")
	}

	if isExsited {

		isBinded := svc.dao.IsBindedUser(userId, userInfo.Id)

		if isBinded {
			return errors.New("用户已绑定")
		}

		err = svc.dao.BindUser(userId, userInfo.Id)
	}

	return err
}

func (svc *Service) UnbindUser(userId, unbindUserId uint32) error {
	isExsited, _ := svc.dao.IsSystemUser("", unbindUserId)

	if !isExsited {
		return errors.New("用户不存在")
	}
	return svc.dao.UnbindUser(userId, unbindUserId)
}

func (svc *Service) GetBindedUserList(userId uint32) ([]model.BindedUserInfo, error) {
	return svc.dao.QueryBindedUserList(userId)
}

func (svc *Service) CheckBindedUser(userId1, userId2 uint32) bool {
	return svc.dao.IsBindedUser(userId1, userId2)
}

func (svc *Service) ChangeAccount(userId uint32) (bool, *model.User) {
	return svc.dao.IsSystemUser("", userId)
}

func (svc *Service) GetUserSetting(userId uint32) (model.UserSetting, error) {
	return svc.dao.QueryUserSetting(userId)
}

func (svc *Service) UpdateUserSetting(req *model.UserSetting) (model.UserSetting, error) {
	return svc.dao.UpdateUserSetting(req)
}

func (svc *Service) IsPrivacyUser(userId uint32) bool {
	bool, user := svc.dao.IsSystemUser("", userId)

	if bool && user.IsPrivacy == 1 {
		return true
	}

	return false
}

func (svc *Service) SaveBindedUserSort(userId uint32, req []model.SaveBindedUserSort) error {
	return svc.dao.SaveBindedUserSort(userId, req)
}
