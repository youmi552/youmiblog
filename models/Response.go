package models

import (
	"blog/config"
	"html/template"
)

//主页返回数据
type HomeResponse struct {
	config.Viewer
	Categorys []Category
	Posts     []PostMore
	Total     int
	Page      int
	Pages     []int
	PageEnd   bool
}

//分类返回数据
type CategoryResponse struct {
	config.Viewer
	CategoryName string
	Categorys    []Category
	Posts        []PostMore
	Total        int
	Page         int
	Pages        []int
	PageEnd      bool
}

//登录返回数据
type LoginResponse struct {
	Token    string   `json:"token"`
	UserInfo UserInfo `json:"userInfo"`
}

//详情返回数据
type DetailResponse struct {
	config.Viewer
	config.SystemConfig
	Article       PostMore
	CategoryNames []string
	Content       template.HTML
}

type WritingResponse struct {
	Title     string
	CdnURL    string
	Categorys []Category
	PostMore  PostMore
}

type PigeonholeResponse struct {
	config.Viewer
	config.SystemConfig
	Categorys []Category
	Lines     map[string][]PostMore
}

type SearchResponse struct {
	Pid   int    `json:"pid"`
	Title string `json:"title"`
}
