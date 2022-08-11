package service

import (
	"blog/config"
	"blog/dao"
	"blog/models"
	"errors"
)

func GetWriting() (models.WritingResponse, error) {
	//获取添加文章页面中展示的数据
	Categorys, err := dao.GetAllCategory()
	if err != nil {
		return models.WritingResponse{}, err
	}
	wr := models.WritingResponse{
		Title:     config.Cfg.Viewer.Title,
		CdnURL:    config.Cfg.System.CdnURL,
		Categorys: Categorys,
	}
	return wr, nil
}
func WritingPost(uid int, postMore models.PostMore, cid int) (models.PostMore, error) {
	//给文章附上注入对应作者
	username, err := dao.GetUserByUid(uid)
	if err != nil {
		return models.PostMore{}, err
	}
	postMore.UserId = uid
	postMore.UserName = username
	//创建文章关联的分类
	var category = []models.Category{{Cid: cid}}
	postMore.Category = category
	//检查数据库受影响的条目数，判断是否成功
	rows, err := dao.CreatePost(postMore)
	if err != nil {
		return models.PostMore{}, err
	}
	if rows == 1 {
		p, err := dao.GetLastPost()
		if err != nil {
			return models.PostMore{}, err
		}
		return p, nil
	}
	return models.PostMore{}, errors.New("添加失败！！！")
}

func GetWritingPost(pid int) (models.WritingResponse, error) {
	//获取修改文章页面展示的数据
	postMore, err := dao.GetPostByPid(pid)
	if err != nil {
		return models.WritingResponse{}, err
	}
	Categorys, err := dao.GetAllCategory()
	if err != nil {
		return models.WritingResponse{}, err
	}
	wr := models.WritingResponse{
		Title:     config.Cfg.Viewer.Title,
		CdnURL:    config.Cfg.System.CdnURL,
		Categorys: Categorys,
		PostMore:  postMore,
	}
	return wr, nil
}
func UpdatePost(postMore models.PostMore, cid int, uid int) (models.PostMore, error) {
	//username, err := dao.GetUserByUid(uid)
	//if err != nil {
	//	return models.PostMore{}, err
	//}
	//postMore.UserId = uid
	//postMore.UserName = username
	Post, err2 := dao.GetPostByPid(postMore.Pid)
	if err2 != nil {
		return models.PostMore{}, err2
	}
	if cid > 0 && cid != Post.Category[0].Cid {
		err := dao.UpdatePostCategory(Post, cid)
		if err != nil {
			return models.PostMore{}, err
		}
	}
	if Post.UserId != uid {
		return models.PostMore{Pid: -1}, errors.New("不可以修改其他人的文章！")
	}
	//检查数据库受影响的条目数，判断是否成功
	rows, err := dao.UpdatePost(postMore)
	if err != nil {
		return models.PostMore{}, err
	}
	if rows == 1 {
		postMore, err = dao.GetPostByPid(postMore.Pid)
		if err != nil {
			return models.PostMore{}, err
		}
		return postMore, nil
	}
	return models.PostMore{}, errors.New("更新失败！！！")
}

func DeletePost(pid int, uid int) error {
	Post, err2 := dao.GetPostByPid(pid)
	if err2 != nil {
		return err2
	}
	if Post.UserId != uid {
		return errors.New("不可以删除其他人的文章！")
	}
	err := dao.DeletePost(pid)
	if err != nil {
		return err
	}
	return nil
}
