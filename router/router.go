package router

import (
	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) {
	//主页
	r.GET("/", Index)
	//分类分页展示
	r.GET("/c/:cid", IndexByCategory)
	//文章详情
	r.GET("/p/:pid", GetDetail)
	//文章归档
	r.GET("/pigeonhole", GetPigeonhole)
	//搜索文章
	r.GET("/search", SearchPost)

	//登录页面
	login := r.Group("login")
	{
		//请求登录页面
		login.GET("/", GetLoginIndex)
		//实现登录
		login.POST("/", Login)
	}

	//文章的修改
	writing := r.Group("/writing")
	{
		//进入文章添加页面
		writing.GET("/", GetWritingIndex)
		//进入对应文章修改页面
		writing.GET("/:pid", GetUpdateIndex)

		writing.POST("/", WritingPost)

		writing.PUT("/", UpdatePost)

		writing.DELETE("/:pid", DeletePost)
	}
	r.POST("/uploadfile", UploadFile)

}
