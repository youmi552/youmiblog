package service

import (
	"blog/config"
	"blog/dao"
	"blog/models"
)

func GetPigeonholeInfo() (models.PigeonholeResponse, error) {
	//获取所有分类信息
	var categorys, err = dao.GetAllCategory()
	if err != nil {
		return models.PigeonholeResponse{}, err
	}
	posts, err := dao.GetAllPost()
	if err != nil {
		return models.PigeonholeResponse{}, err
	}
	lines := make(map[string][]models.PostMore)
	for _, post := range posts {
		at := post.CreatedAt
		month := at.Format("2006-01")
		lines[month] = append(lines[month], post)
	}

	var pr = models.PigeonholeResponse{
		Viewer:       config.Cfg.Viewer,
		SystemConfig: config.Cfg.System,
		Categorys:    categorys,
		Lines:        lines,
	}
	return pr, nil
}
