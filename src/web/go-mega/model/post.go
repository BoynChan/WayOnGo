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

// 根据用户id获取发表过的文章
// 并根据用户参数进行分页
func GetPostsByUserIDPageAndLimit(id, page, limit int) (*[]Post, int, error) {
	var total int
	var posts []Post
	offset := (page - 1) * limit
	if err := db.Preload("User").
		Order("timestamp desc").
		Where("user_id=?", id).
		Offset(offset).
		Limit(limit).
		Find(&posts).Error; err != nil {
		return nil, total, err
	}
	db.Model(&Post{}).Where("user_id=?", id).Count(&total)
	return &posts, total, nil
}

// 根据时间获取所有文章
func GetPostsByPageAndLimit(page, limit int) (*[]Post, int, error) {
	var total int
	var posts []Post
	offset := (page - 1) * limit
	if err := db.Preload("User").
		Order("timestamp desc").
		Offset(offset).
		Limit(limit).
		Find(&posts).Error; err != nil {
		return nil, total, err
	}
	db.Model(&Post{}).Count(&total)
	return &posts, total, nil
}

// 根据用户id获取发表过的文章
func GetPostsByUserID(id int) (*[]Post, error) {
	var posts []Post
	//Preload相当于一个Join Table
	if err := db.Preload("User").Where("user_id=?", id).Find(&posts).Error; err != nil {
		return nil, err
	}
	return &posts, nil
}
