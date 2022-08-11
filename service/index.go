package service

import (
	"blog/config"
	"blog/dao"
	"blog/models"
	"blog/util"
)

func GetAllIndexInfo(pagenumber int, pagesize int) (*models.HomeResponse, error) {
	//获取所有分类信息
	categorys, err := dao.GetAllCategory()
	if err != nil {
		return nil, err
	}
	//获取所有文章信息并且分类
	posts, err := dao.GetPostPage(pagenumber, pagesize)
	if err != nil {
		return nil, err
	}
	//处理掉超文本中的标签
	for i := range posts {
		posts[i].Content = util.TrimHtml(posts[i].Content)
	}
	//计录所有文章数
	var total = dao.CountGetAllPost()
	//计算分页总数
	pagesCount := (total-1)/pagesize + 1
	//添加所有的分页下标用于前端分页标签
	var pages []int
	for i := 0; i < pagesCount; i++ {
		pages = append(pages, i+1)
	}
	//封装好数据传给控制层
	var hr = &models.HomeResponse{
		config.Cfg.Viewer,
		categorys,
		posts,
		total,
		pagenumber,
		pages,
		pagenumber != pagesCount, //若没有到达分页页数底部，则返回false
	}
	return hr, nil
}
