package function

import "blog/models"

func CreateSuccess(data interface{}) models.Result {
	var result = models.Result{
		Msg:  "添加成功！",
		Code: 200,
		Data: data,
	}
	return result
}
func UpdateSuccess(data interface{}) models.Result {
	var result = models.Result{
		Msg:  "更新成功！",
		Code: 200,
		Data: data,
	}
	return result
}
func TokenTransferError() models.Result {
	var result = models.Result{
		Msg:  "登录已过期！",
		Code: 1003,
		Data: nil,
	}
	return result
}
func NoRootDelete(err error) models.Result {
	var result = models.Result{
		Msg:  "不可以删除悠米的文章！！",
		Code: 1010,
		Data: err,
	}
	return result
}
func NoRootUpdate(err error) models.Result {
	var result = models.Result{
		Msg:  "不可以修改其他人的文章！！",
		Code: 1011,
		Data: err,
	}
	return result
}
func NoRoot(err error) models.Result {
	var result = models.Result{
		Msg:  "没有权限！请登录！",
		Code: 1006,
		Data: err,
	}
	return result
}
func DeleteSuccess() models.Result {
	var result = models.Result{
		Msg:  "删除成功",
		Code: 200,
		Data: nil,
	}
	return result
}
