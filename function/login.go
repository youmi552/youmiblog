package function

import (
	"blog/models"
)

func LoginSuccess(data interface{}) models.Result {
	var result = models.Result{
		Msg:  "登陆成功",
		Code: 200,
		Data: data,
	}
	return result
}
func UserNameIsNull() models.Result {
	var result = models.Result{
		Msg:  "用户名不能为空",
		Code: 1000,
		Data: nil,
	}
	return result
}

func PassWordIsNull() models.Result {
	var result = models.Result{
		Msg:  "密码不能为空",
		Code: 1001,
		Data: nil,
	}
	return result
}

func UserNameOrPasswordError() models.Result {
	var result = models.Result{
		Msg:  "用户名或密码错误",
		Code: 1002,
		Data: nil,
	}
	return result
}
