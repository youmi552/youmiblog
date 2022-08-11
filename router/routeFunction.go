package router

import (
	"blog/config"
	"blog/function"
	"blog/models"
	"blog/service"
	"blog/util"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
	"path"
	"strconv"
)

func Index(c *gin.Context) {
	//获取请求的中的页数信息并转化为int类型
	pagenumber, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	//处理可能遇到的转换类型错误
	if err != nil {
		log.Println("转换类型失败！:")
		log.Println(err)
		c.HTML(http.StatusBadRequest, "index.html", nil)
		return
	}
	//检查pagenumber的合法性
	pagenumber = util.CheckPageNumber(pagenumber)
	//设置每页显示的数据
	pagesize := 10
	//从业务层获取主页数据
	hr, err := service.GetAllIndexInfo(pagenumber, pagesize)
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusBadRequest, "index", nil)
		return
	}
	//返回给前端
	c.HTML(http.StatusOK, "index.html", hr)
}

func IndexByCategory(c *gin.Context) {
	//获取请求中的分类id
	cid, err := strconv.Atoi(c.Param("cid"))
	//处理可能遇到的转换类型错误
	if err != nil {
		log.Println("转换类型失败！:")
		log.Println(err)
		c.HTML(http.StatusBadRequest, "category.html", nil)
		return
	}
	//获取请求的中的页数信息并转化为int类型
	pagenumber, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	//处理可能遇到的转换类型错误
	if err != nil {
		log.Println("转换类型失败！:")
		log.Println(err)
		c.HTML(http.StatusBadRequest, "category.html", nil)
		return
	}
	pagenumber = util.CheckPageNumber(pagenumber)
	//设置每页显示的数据
	pagesize := 10
	//从业务层获取分类页中的数据
	cr, err := service.GetAllCategoryInfo(cid, pagenumber, pagesize)
	if err != nil {
		c.HTML(http.StatusBadRequest, "category.html", err)
		return
	}
	//返回给前端
	c.HTML(http.StatusOK, "category.html", cr)
}

func GetPigeonhole(c *gin.Context) {
	pr, err := service.GetPigeonholeInfo()
	if err != nil {
		c.HTML(http.StatusBadRequest, "pigeonhole.html", err)
		return
	}
	c.HTML(http.StatusOK, "pigeonhole.html", pr)
}

func SearchPost(c *gin.Context) {
	condition := c.DefaultQuery("val", "")
	sp, err := service.GetSearchInfo(condition)
	if err != nil {
		result := function.Fail(err)
		c.JSON(http.StatusOK, result)
		return
	}
	result := function.Success(sp)
	c.JSON(http.StatusOK, result)
}

func GetLoginIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", config.Cfg.Viewer)
}

func Login(c *gin.Context) {
	//接收前端传来的json数据
	var user models.User
	c.ShouldBind(&user)
	//如果用户名为空,则向前端返回错误数据
	if user.UserName == "" {
		c.JSON(http.StatusOK, function.UserNameIsNull())
		return
	}
	//如果密码为空,则向前端返回错误数据
	if user.Passwd == "" {
		c.JSON(http.StatusOK, function.PassWordIsNull())
		return
	}
	//将用户信息提交给业务层
	lr, err := service.Login(user)
	if err != nil {
		result := function.Fail(err)
		c.JSON(http.StatusOK, result)
		return
	}
	//将登录成功数据传给前端
	result := function.LoginSuccess(lr)
	c.JSON(http.StatusOK, result)
}

func GetDetail(c *gin.Context) {
	//获取请求中的文章的id
	pid, err := strconv.Atoi(c.Param("pid"))
	//处理可能遇到的转换类型错误
	if err != nil {
		log.Println("转换类型失败！:")
		log.Println(err)
		c.HTML(http.StatusBadRequest, "detail.html", nil)
		return
	}
	//从业务层获取文章详情页面中的数据
	dr, err := service.GetDetailInfo(pid)
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusBadRequest, "detail.html", nil)
		return
	}
	//返回给前端
	c.HTML(http.StatusOK, "detail.html", dr)
}

func GetWritingIndex(c *gin.Context) {
	wr, err := service.GetWriting()
	if err != nil {
		c.HTML(http.StatusBadRequest, "writing.html", nil)
	}
	c.HTML(http.StatusOK, "writing.html", wr)
}

func GetUpdateIndex(c *gin.Context) {
	//获取请求中的分类id
	pid, err := strconv.Atoi(c.Param("pid"))
	//处理可能遇到的转换类型错误
	if err != nil {
		log.Println("转换类型失败！:")
		log.Println(err)
		result := function.Fail(err)
		c.JSON(http.StatusOK, result)
		return
	}
	//如果文章id>0则查找对应文章数据显示在页面上
	if pid > 0 {
		wr, err := service.GetWritingPost(pid)
		if err != nil {
			result := function.Fail(err)
			c.JSON(http.StatusOK, result)
		}
		result := function.Success(wr)
		c.JSON(http.StatusOK, result)
		return
	}
	c.JSON(http.StatusOK, function.Fail(errors.New("文章不存在！")))
}

func WritingPost(c *gin.Context) {
	//获取token
	token := c.Request.Header.Get("Authorization")
	//解密token获得用户uid
	_, claim, err := util.ParseToken(token)
	//解密失败则返回登录已过期提示
	if err != nil {
		result := function.TokenTransferError()
		c.JSON(http.StatusOK, result)
		return
	}
	//接收前端传来的json数据
	var postMore models.PostMore
	//c.ShouldBind(&postMore)
	c.ShouldBindBodyWith(&postMore, binding.JSON)
	var category models.CategoryId
	c.ShouldBindBodyWith(&category, binding.JSON)
	cid, err := strconv.Atoi(category.Cid)
	if err != nil {
		log.Println("转换类型失败！:")
		log.Println(err)
		result := function.Fail(err)
		c.JSON(http.StatusOK, result)
	}
	//进行添加文章操作
	p, err := service.WritingPost(claim.Uid, postMore, cid)
	if err != nil {
		result := function.Fail(err)
		c.JSON(http.StatusOK, result)
		return
	}
	result := function.CreateSuccess(p)
	c.JSON(http.StatusOK, result)
}

func UpdatePost(c *gin.Context) {
	//获取token
	token := c.Request.Header.Get("Authorization")
	//解密token获得用户uid
	_, claim, err := util.ParseToken(token)
	//解密失败则返回登录已过期提示
	if err != nil || claim.Uid < 1 {
		result := function.NoRoot(err)
		c.JSON(http.StatusOK, result)
		return
	}
	//接收前端传来的json数据
	var postMore models.PostMore
	//c.ShouldBind(&postMore)
	c.ShouldBindBodyWith(&postMore, binding.JSON)
	var categoryId models.Category
	c.ShouldBindBodyWith(&categoryId, binding.JSON)
	//cid, err := strconv.Atoi(categoryId.Cid)
	//fmt.Println("categoryid:", categoryId.Cid)
	//return
	if err != nil {
		log.Println("转换类型失败！:")
		log.Println(err)
		result := function.Fail(err)
		c.JSON(http.StatusOK, result)
	}
	//进行修改文章操作
	p, err := service.UpdatePost(postMore, categoryId.Cid, claim.Uid)
	if err != nil {
		if p.Pid == -1 {
			result := function.NoRootUpdate(err)
			c.JSON(http.StatusOK, result)
			return
		}
		result := function.Fail(err)
		c.JSON(http.StatusOK, result)
		return
	}
	result := function.UpdateSuccess(p)
	c.JSON(http.StatusOK, result)
}
func DeletePost(c *gin.Context) {
	//获取token
	token := c.Request.Header.Get("Authorization")
	//解密token获得用户uid
	_, claim, err := util.ParseToken(token)
	//解密失败则返回登录已过期提示
	if err != nil || claim.Uid < 1 {
		result := function.NoRoot(err)
		c.JSON(http.StatusOK, result)
		return
	}
	//获取请求中的分类id
	pid, err := strconv.Atoi(c.Param("pid"))
	//处理可能遇到的转换类型错误
	if err != nil {
		log.Println("转换类型失败！:")
		log.Println(err)
		result := function.Fail(err)
		c.JSON(http.StatusOK, result)
		return
	}
	err = service.DeletePost(pid, claim.Uid)
	if err != nil {
		if claim.Uid != 1 {
			result := function.NoRootDelete(err)
			c.JSON(http.StatusOK, result)
			return
		}
		result := function.Fail(err)
		c.JSON(http.StatusOK, result)
		return
	}
	result := function.DeleteSuccess()
	c.JSON(http.StatusOK, result)

}
func UploadFile(c *gin.Context) {
	//读取文件
	f, err := c.FormFile("f1")
	//设置文件上传错误的返回值
	if err != nil {
		fmt.Println("上传文件失败")
		result := function.UploadError()
		c.JSON(http.StatusOK, result)
		return
	}
	//dst := fmt.Sprintf("./%s", f.Filename)
	//设置上传路径
	dst := path.Join("./file", f.Filename)
	c.SaveUploadedFile(f, dst)
	//设置映射路径并返回的前端
	url := fmt.Sprintf("https://5706a5m231.zicp.fun/%s", dst)
	result := function.UploadSuccess(url)
	c.JSON(http.StatusOK, result)

}
