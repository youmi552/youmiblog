package models

import "time"

//用户
type User struct {
	Uid       int       `gorm:"primarykey" column:"uid"`
	UserName  string    `form:"username" json:"username"`
	Passwd    string    `form:"passwd" json:"passwd"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"createAt"` // 创建时间
	UpdatedAt time.Time `json:"updateAt"` // 更新时间
}

//用户信息
type UserInfo struct {
	Uid      int    `json:"uid"`
	UserName string `json:"username"`
	Avatar   string `json:"avatar"`
}
