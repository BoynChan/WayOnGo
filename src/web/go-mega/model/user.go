package model

import (
	"fmt"
	"time"
)

// Author:Boyn
// Date:2020/2/23

// 表示用户
type User struct {
	ID           int    `gorm:"primary_key"`
	Username     string `gorm:"type:varchar(255)"`
	Email        string `gorm:"type:varchar(255)"`
	PasswordHash string `gorm:"type:varchar(255)"`
	Posts        []Post
	Followers    []*User `gorm:"many2many:follower;association_jointable_foreignkey:follower_id"`
	LastSeen     *time.Time
	AboutMe      string `gorm:"type:varchar(255)"`
	Avatar       string `gorm:"type:varchar(255)"`
}

func (u *User) SetAvatar(email string) {
	u.Avatar = fmt.Sprintf("https://www.gravatar.com/avatar/%s?d=identicon", Md5(email))
}

func (u *User) SetPassword(pwd string) {
	u.PasswordHash = GeneratePasswordHash(pwd)
}

func (u *User) CheckPassword(pwd string) bool {
	return GeneratePasswordHash(pwd) == u.PasswordHash
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	if err := db.Where("username=?", username).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// 向表中添加一个User
func AddUser(username, password, email string) error {
	user := User{
		Username: username,
		Email:    email,
	}
	user.SetAvatar(email)
	user.SetPassword(password)
	return db.Create(&user).Error
}

func UpdateUserByUsername(username string, contents map[string]interface{}) error {
	item, err := GetUserByUsername(username)
	if err != nil {
		return err
	}
	return db.Model(item).Updates(contents).Error
}

func UpdateLastSeen(username string) error {
	contents := map[string]interface{}{"last_seen": time.Now()}
	return UpdateUserByUsername(username, contents)
}

func UpdateAboutMe(username, aboutMe string) error {
	contents := map[string]interface{}{"about_me": aboutMe}
	return UpdateUserByUsername(username, contents)
}
