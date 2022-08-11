package service

import (
	"blog/dao"
	"blog/models"
	"blog/util"
	"errors"
)

func Login(user models.User) (models.LoginResponse, error) {
	//将用户密码添加盐值并进行Md5加密
	user.Passwd = util.Md5Crypt(user.Passwd, "youmi")
	//查询是否存在该用户
	user2, err := dao.SelectUserByUserNameAndPasswd(user)
	//如果未能查到数据则添加错误信息到result并报错
	if user2.UserName == "" || user2.Passwd == "" {
		err := errors.New("用户名或密码错误")
		return models.LoginResponse{}, err
	}
	//UserInfo中注入属性
	var userInfo = models.UserInfo{
		Uid:      user2.Uid,
		UserName: user2.UserName,
		Avatar:   user2.Avatar,
	}
	//生成token
	token, err := util.Award(&userInfo.Uid)
	if err != nil {
		return models.LoginResponse{}, err
	}
	//整理数据到lr
	lr := models.LoginResponse{
		UserInfo: userInfo,
		Token:    token,
	}
	//将数据整合到result中
	return lr, nil
}
