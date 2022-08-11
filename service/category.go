package service

import (
	"blog/config"
	"blog/dao"
	"blog/models"
	"blog/util"
)

func GetAllCategoryInfo(cid int, pagenumber int, pagesize int) (*models.CategoryResponse, error) {
	//获取所有分类信息
	categorys, err := dao.GetAllCategory()
	if err != nil {
		return nil, err
	}
	//获取分类名称
	categoryName, err := dao.GetCategoryName(cid)
	if err != nil {
		return nil, err
	}
	//根据文章分类获取文章信息并分页
	posts, err := dao.GetPostPageByCid(cid, pagenumber, pagesize)
	//去掉标签展示数据
	for i := range posts {
		posts[i].Content = util.TrimHtml(posts[i].Content)
	}
	//记录该该类型文章总数
	total := dao.CountGetPostByCid(cid)
	//计算分页总数
	pagesCount := (total-1)/pagesize + 1
	//添加所有的分页下标用于前端分页标签
	var pages []int
	for i := 0; i < pagesCount; i++ {
		pages = append(pages, i+1)
	}
	//封装好数据传给控制层
	var cr = &models.CategoryResponse{
		config.Cfg.Viewer,
		categoryName,
		categorys,
		posts,
		total,
		pagenumber,
		pages,
		pagenumber != pagesCount, //若没有到达分页页数底部，则返回false
	}
	return cr, nil
}
