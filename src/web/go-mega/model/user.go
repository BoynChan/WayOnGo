package model

// Author:Boyn
// Date:2020/2/23

// 表示用户
type User struct {
	ID           int    `gorm:"primary_key"`
	Username     string `gorm:"type:varchar(255)"`
	Email        string `gorm:"type:varchar(255)"`
	PasswordHash string `gorm:"type:varchar(255)"`
	Posts        []Post
	Followers    []*User `gorm:many2many:follower;association_jointable_foreignkey:follower_id`
}
