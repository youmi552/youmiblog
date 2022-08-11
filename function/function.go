package function

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"time"
)

func Function(r *gin.Engine) {
	r.SetFuncMap(template.FuncMap{
		"isODD": func(num int) bool {
			return num%2 == 0
		},
		"getNextName": func(strs []string, index int) string {
			return strs[index+1]
		},
		"date": func(layout string) string {
			return time.Now().Format(layout)
		},
		"dateDay": func(date time.Time) string {
			return date.Format("2006-01-02 15:04:05")
		},
	})
}
