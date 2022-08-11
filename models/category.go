package models

import (
	"time"
)

type Category struct {
	Cid       int        `gorm:"primarykey" column:"cid" json:"categoryId"`
	Name      string     `json:"categoryName"`
	CreatedAt time.Time  `json:"createAt"` // 创建时间
	UpdatedAt time.Time  `json:"updateAt"` // 更新时间
	PostMore  []PostMore `gorm:"many2many:post_more_category;"`
}
type CategoryId struct {
	Cid string `json:"categoryId"`
}
