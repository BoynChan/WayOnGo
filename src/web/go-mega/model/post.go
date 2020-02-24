package model

import "time"

// Author:Boyn
// Date:2020/2/23

// 表示文章结构
type Post struct {
	ID        int `gorm:"primary_key"`
	UserID    int
	User      User
	Body      string     `gorm:"type:varchar(255)"`
	Timestamp *time.Time `sql:"DEFAULT:current_timestamp"`
}
