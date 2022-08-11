package dao

import (
	"blog/models"
)

func SelectUserByUserNameAndPasswd(user models.User) (models.User, error) {
	var user2 = models.User{
		UserName: "",
		Passwd:   "",
	}
	DB := db.Where("user_name=? And passwd= ?", user.UserName, user.Passwd).First(&user2)
	if DB.Error != nil {
		return models.User{}, DB.Error
	}
	return user2, nil
}
func GetUserByUid(uid int) (string, error) {
	var username string
	DB := db.Model(&models.User{}).Where("uid=?", uid).Select("user_name").First(&username)
	if DB.Error != nil {
		return "", DB.Error
	}
	return username, nil
}
