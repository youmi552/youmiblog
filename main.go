package main

import (
	"blog/function"
	"blog/router"
	"github.com/gin-gonic/gin"
	"image"
	"image/jpeg"
)

func init() {
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
}
func main() {
	r := gin.Default()
	function.Function(r)                       //加载模板函数
	r.Static("/resource", "./public/resource") //处理静态资源路径
	r.Static("/file", "./file")
	r.LoadHTMLGlob("templates/**/*") //扫描html所有页面
	//r.LoadHTMLFiles("./template/home.html",
	//	"./template/index.html",
	//	"./template/layout/footer.html",
	//	"./template/layout/header.html",
	//	"./template/layout/pagination.html",
	//	"./template/layout/personal.html",
	//	"./template/layout/post-list.html")
	router.Route(r)
	r.Run(":80")
}
