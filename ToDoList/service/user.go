package service

import (
	"main/ToDoList/model"
	"main/ToDoList/pkg/utils"
	"main/ToDoList/serializer"

	"github.com/jinzhu/gorm"
)

type UserService struct {
	UserName string `form:"user_name" json:"user_name" `
	Password string `form:"password" json:"password"`
}

func (service *UserService) Register() serializer.Response {
	var user model.User
	var count int
	model.DB.Model(&model.User{}).Where("user_name=?", service.UserName).First(&user).Count(&count)
	if count == 1 {
		return serializer.Response{
			Status: 400,
			Msg:    " 用户已存在",
		}
	}
	user.UserName = service.UserName
	//密码加密
	if err := user.SetPassword(service.Password); err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    err.Error(),
		}
	}

	// 创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "数据库操作错误",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "用户注册成功了",
	}

}

func (service *UserService) Login() serializer.Response {
	var user model.User

	//查找数据库中是否存在该用户
	if err := model.DB.Where("user_name=?", service.UserName).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return serializer.Response{
				Status: 400,
				Msg:    "用户不存在, 请先注册",
			}
		}

		//如果是其他错误
		return serializer.Response{
			Status: 500,
			Msg:    "数据库错误",
		}
	}
	if !user.CheckPassword(service.Password) {
		return serializer.Response{
			Status: 400,
			Msg:    "密码错误",
		}
	}
	//发送token，为了其他功能需要身份验证所给前端存储
	token, err := utils.GenerateToken(user.ID, service.UserName, service.Password)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "Token签发错误",
		}
	}
	return serializer.Response{
		Status: 200,
		Data: serializer.TokenData{
			User:  serializer.BuildUser(user),
			Token: token,
		},
		Msg: "登录成功",
	}
}
