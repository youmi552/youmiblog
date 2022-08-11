package function

import (
	"blog/models"
	"net/http"
)

func Success(data interface{}) models.Result {
	var result = models.Result{
		Msg:  "请求成功！",
		Code: http.StatusOK,
		Data: data,
	}
	return result
}

func Fail(err error) models.Result {
	var result = models.Result{
		Msg:  "请求失败！",
		Code: http.StatusBadRequest,
		Data: err,
	}
	return result
}
