package myerror

import (
	"github.com/gin-gonic/gin"
	"log"
)

func StrconvError(c *gin.Context, err error) {
	if err != nil {
		log.Println("转换类型失败！:")
		log.Println(err)
		c.Error(err)
	}
}

func TokenError(err error) {
	if err != nil {
		log.Println("token未生成！:", err)
	}
}
