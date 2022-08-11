package dao

import (
	"blog/models"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

//获取所有文章并分页
func GetPostPage(pagenumber int, pagesize int) ([]models.PostMore, error) {
	var PostMores []models.PostMore
	DB := db.Preload("Category").Limit(pagesize).Offset((pagenumber - 1) * pagesize).Find(&PostMores) //在关联查询的基础上预加载Game分页查询
	if DB.Error != nil {
		return nil, DB.Error
	}
	return PostMores, nil
}

//计算所有文章数
func CountGetAllPost() int {
	var count int64
	db.Model(&models.PostMore{}).Count(&count)
	return int(count)
}

//根据类别id查询所有文章
func GetPostPageByCid(cid int, pagenumber int, pagesize int) ([]models.PostMore, error) {
	c := models.Category{Cid: cid}
	var PostMores []models.PostMore
	err := db.Model(&c).Limit(pagesize).Offset((pagenumber - 1) * pagesize).Preload("Category").Association("PostMore").Find(&PostMores)
	if err != nil {
		return nil, err
	}
	return PostMores, nil
}

//根据类别id查询文章数量
func CountGetPostByCid(cid int) int {
	c := models.Category{Cid: cid}
	var PostMores []models.PostMore
	db.Model(&c).Preload("Category").Association("PostMore").Find(&PostMores)
	count := 0
	for i := range PostMores {
		if PostMores[i].Pid != 0 {
			count++
		}
	}
	return count
}

//根据文章id查询文章内容
func GetPostByPid(pid int) (models.PostMore, error) {
	var PostMore models.PostMore
	DB := db.Where("pid=?", pid).Preload("Category").First(&PostMore)
	if DB.Error != nil {
		return models.PostMore{}, DB.Error
	}
	return PostMore, nil
}

//创建文章
func CreatePost(postMore models.PostMore) (int, error) {
	DB := db.Create(&postMore)
	if DB.Error != nil {
		return 0, DB.Error
	}
	return int(DB.RowsAffected), nil
}

//获取最后的文章
func GetLastPost() (models.PostMore, error) {
	var postMore models.PostMore
	DB := db.Last(&postMore)
	if DB.Error != nil {
		return models.PostMore{}, DB.Error
	}
	return postMore, nil
}

//更新文章数据
func UpdatePost(postMore models.PostMore) (int, error) {
	DB := db.Model(&models.PostMore{}).Where("pid=?", postMore.Pid).Updates(&postMore)
	if DB.Error != nil {
		return 0, DB.Error
	}
	return int(DB.RowsAffected), nil
}
func UpdatePostCategory(p models.PostMore, cid2 int) error {
	//c1 := models.Category{Cid: cid1}
	c2 := models.Category{Cid: cid2}
	err := db.Model(&p).Association("Category").Clear()
	err = db.Model(&p).Association("Category").Append(&c2)
	if err != nil {
		return err
	} //替换所有关联
	return nil
}
func GetAllPost() ([]models.PostMore, error) {
	var p []models.PostMore
	DB := db.Find(&p)
	if DB.Error != nil {
		return nil, DB.Error
	}
	return p, nil
}

func GetPostByCondition(condition string) ([]models.SearchResponse, error) {
	var sp []models.SearchResponse
	DB := db.Model(&models.PostMore{}).Where("title like ?", "%"+condition+"%").Find(&sp)
	if DB.Error != nil {
		return nil, DB.Error
	}
	return sp, nil
}

func DeletePost(pid int) error {
	var p = models.PostMore{Pid: pid}
	DB := db.Delete(&p)
	if DB.Error != nil {
		return DB.Error
	}
	return nil
}

func PostByIdExists(pid int) bool {
	defer rd.Close()
	user := fmt.Sprintf("user:%s", pid)
	bool, err := redis.Bool(rd.Do("hexist", user))
	if err != nil {
		log.Println(err)
	}
	return bool
}

//func GetPostByRedis(pid int) (models.PostMore, error) {
//	defer rd.Close()
//	user := fmt.Sprintf("user:%s", pid)
//	pid2 := fmt.Sprintf("%s:pid", user)
//	title := fmt.Sprintf("%s:title", user)
//	slug := fmt.Sprintf("%s:slug", user)
//	content := fmt.Sprintf("%s:content", user)
//	markdown := fmt.Sprintf("%s:markdown", user)
//	userId := fmt.Sprintf("%s:userId", user)
//	userName := fmt.Sprintf("%s:userName", user)
//	viewCount := fmt.Sprintf("%s:viewCount", user)
//	types:= fmt.Sprintf("%s:type", user)
//	createdAt := fmt.Sprintf("%s:createdAt", user)
//	updatedAt := fmt.Sprintf("%s:updatedAt", user)
//	var posts =models.PostMore{
//		Pid:redis.Int(rd.Do(ge)),
//			}
//}
