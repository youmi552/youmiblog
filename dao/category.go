package dao

import "blog/models"

func GetAllCategory() ([]models.Category, error) { //查询所有的分类
	var categorys []models.Category
	DB := db.Find(&categorys)
	if DB.Error != nil {
		return nil, DB.Error
	}
	return categorys, nil
}
func GetCategoryName(cid int) (string, error) {
	var categoryName string
	DB := db.Model(&models.Category{}).Select("name").Where("cid=?", cid).First(&categoryName)
	if DB.Error != nil {
		return "", DB.Error
	}
	return categoryName, nil
}
func GetCategoryNameByPid(pid int) ([]string, error) {
	var p = models.PostMore{Pid: pid}
	var categoryName []string
	err := db.Model(&p).Select("name").Preload("PostMore").Association("Category").Find(&categoryName)
	if err != nil {
		return nil, err
	}
	return categoryName, nil
}
