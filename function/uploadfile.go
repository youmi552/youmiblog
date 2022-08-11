package function

import "blog/models"

func UploadSuccess(data interface{}) models.Result {
	var result = models.Result{
		Msg:  "上传成功！",
		Code: 200,
		Data: data,
	}
	return result
}

func UploadError() models.Result {
	var result = models.Result{
		Msg:  "上传文件失败！",
		Code: 1004,
		Data: nil,
	}
	return result
}
