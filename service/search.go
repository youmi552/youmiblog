package service

import (
	"blog/dao"
	"blog/models"
)

func GetSearchInfo(condition string) ([]models.SearchResponse, error) {
	sp, err := dao.GetPostByCondition(condition)
	if err != nil {
		return nil, err
	}
	return sp, nil
}
