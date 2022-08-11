package service

import (
	"blog/config"
	"blog/dao"
	"blog/models"
	"html/template"
)

func GetDetailInfo(pid int) (*models.DetailResponse, error) {
	//获取分类名称
	categoryNames, err := dao.GetCategoryNameByPid(pid)
	if err != nil {
		return nil, err
	}
	//bool := dao.PostByIdExists(pid)
	//if bool {
	//	posts, err := dao.GetPostByRedis(pid)
	//}
	//根据文章分类获取文章信息并分页
	posts, err := dao.GetPostByPid(pid)
	if err != nil {
		return nil, err
	}
	posts.ViewCount++
	//封装好数据传给控制层
	var dr = &models.DetailResponse{
		config.Cfg.Viewer,
		config.Cfg.System,
		posts,
		categoryNames,
		template.HTML(posts.Content),
	}
	_, err = dao.UpdatePost(posts)
	if err != nil {
		return nil, err
	}
	return dr, nil
}
