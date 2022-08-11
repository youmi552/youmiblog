package main

import (
	"blog/dao"
	"fmt"
)

func main() {
	//fmt.Println(dao.GetAllCategory())
	//fmt.Println(dao.GetPostPage(3, 10))
	//fmt.Println(dao.CountGetAllPost())
	//fmt.Println(dao.GetPostPageByCid(1, 1, 10)
	//fmt.Println(dao.CountGetPostByCid(1))
	//fmt.Println(dao.GetCategoryName(1))
	//fmt.Println(dao.GetCategoryNameByPid(1))
	fmt.Println(dao.GetPostByPid(1))
	//fmt.Println(dao.GetUserByUid(2))
}
