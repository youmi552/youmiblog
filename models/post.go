package models

import (
	"html/template"
	"time"
)

type PostMore struct {
	Pid       int        `gorm:"primarykey" column:"pid" json:"pid"` // 文章ID
	Title     string     `json:"title"`                              // 文章ID
	Slug      string     `json:"slug"`                               // 自定也页面 path
	Content   string     `gorm:"type:longtext" json:"content"`       // 文章的html
	Markdown  string     `gorm:"type:longtext" json:"markdown"`      // 文章的Markdown
	UserId    int        `json:"userId"`                             // 用户id
	UserName  string     `json:"userName"`                           // 用户名
	ViewCount int        `json:"viewCount"`                          // 查看次数
	Type      int        `json:"type"`                               // 文章类型 0 普通，1 自定义文章
	CreatedAt time.Time  `json:"createAt"`
	UpdatedAt time.Time  `json:"updateAt"`
	Category  []Category `gorm:"many2many:post_more_category;"`
}

type PostReq struct {
	Pid        int    `json:"pid"`
	Title      string `json:"title"`
	Slug       string `json:"slug"`
	Content    string `json:"content"`
	Markdown   string `json:"markdown"`
	CategoryId int    `json:"categoryId"`
	UserId     int    `json:"userId"`
	Type       int    `json:"type"`
}

type SearchResp struct {
	Pid   int    `orm:"pid" json:"pid"` // 文章ID
	Title string `orm:"title" json:"title"`
}

type Writing struct {
	PostMore PostMore
	Category Category
}
type PostMoreUpdate struct {
	Pid       int           `gorm:"primarykey" column:"pid" json:"pid"` // 文章ID
	Title     string        `json:"title"`                              // 文章ID
	Slug      string        `json:"slug"`                               // 自定也页面 path
	Content   template.HTML `gorm:"type:longtext" json:"content"`       // 文章的html
	Markdown  string        `gorm:"type:longtext" json:"markdown"`      // 文章的Markdown
	UserId    int           `json:"userId"`                             // 用户id
	UserName  string        `json:"userName"`                           // 用户名
	ViewCount int           `json:"viewCount"`                          // 查看次数
	Type      int           `json:"type"`                               // 文章类型 0 普通，1 自定义文章
	CreatedAt time.Time     `json:"createAt"`
	UpdatedAt time.Time     `json:"updateAt"`
	Category  []Category    `gorm:"many2many:post_more_category;"`
}

type Pid struct {
	Pid int `json:"id"`
}
