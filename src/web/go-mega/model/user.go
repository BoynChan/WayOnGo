package model

import (
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
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

// 设置用户头像
// 使用gravatar的随机头像功能,将email地址md5化之后引用这个网站的头像,因为md5值固定
// 所以头像不变
func (u *User) SetAvatar(email string) {
	u.Avatar = fmt.Sprintf("https://www.gravatar.com/avatar/%s?d=identicon", Md5(email))
}

// 设置用户密码
func (u *User) SetPassword(pwd string) {
	u.PasswordHash = GeneratePasswordHash(pwd)
}

// 检查用户密码是否正确
func (u *User) CheckPassword(pwd string) bool {
	return GeneratePasswordHash(pwd) == u.PasswordHash
}

// 根据username从数据库中取出用户
func GetUserByUsername(username string) (*User, error) {
	var user User
	if err := db.Where("username=?", username).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// 根据email从数据库中取出用户
func GetUserByEmail(email string) (*User, error) {
	var user User
	if err := db.Where("email=?", email).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// 向表中添加一个User
// 在添加新用户时,将自己加入到自己的关注列表
func AddUser(username, password, email string) error {
	user := User{
		Username: username,
		Email:    email,
	}
	user.SetAvatar(email)
	user.SetPassword(password)
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return user.FollowSelf()
}

// 更新数据库中某用户的信息
func UpdateUserInfoByUsername(username string, contents map[string]interface{}) error {
	item, err := GetUserByUsername(username)
	if err != nil {
		return err
	}
	return db.Model(item).Updates(contents).Error
}

// 更新最后操作信息
func UpdateLastSeen(username string) error {
	contents := map[string]interface{}{"last_seen": time.Now()}
	return UpdateUserInfoByUsername(username, contents)
}

// 更新 "关于我" 信息
func UpdateAboutMe(username, aboutMe string) error {
	contents := map[string]interface{}{"about_me": aboutMe}
	return UpdateUserInfoByUsername(username, contents)
}

// 更新密码
func UpdatePassword(username, password string) error {
	contents := map[string]interface{}{"password": Md5(password)}
	return UpdateUserInfoByUsername(username, contents)
}

// 关注某人
// 这是一个User的方法,表示本结构作为关注人,传入参数作为被关注人
func (u *User) Follow(username string) error {
	other, err := GetUserByUsername(username)
	if err != nil {
		return err
	}
	return db.Model(other).Association("Followers").Append(u).Error
}

// 取消关注某人
// 这是一个User的方法,表示本结构作为关注人,传入参数作为被关注人
func (u *User) Unfollow(username string) error {
	other, err := GetUserByUsername(username)
	if err != nil {
		return err
	}
	return db.Model(other).Association("Followers").Delete(u).Error
}

// 关注自己
func (u *User) FollowSelf() error {
	return db.Model(u).Association("Followers").Append(u).Error
}

// 统计粉丝数量
func (u *User) FollowersCount() int {
	return db.Model(u).Association("Followers").Count()
}

// 返回自己关注的人的ID合集
func (u *User) FollowingIDs() []int {
	var ids []int
	rows, err := db.Table("follower").Where("follower_id = ?", u.ID).Select("user_id,follower_id").Rows()
	if err != nil {
		log.Println("[FollowingIDs] 统计关注者ID集合出错:", err)
	}
	defer rows.Close()
	for rows.Next() {
		var id, followerID int
		rows.Scan(&id, &followerID)
		ids = append(ids, id)
	}
	return ids
}

// 统计自己关注的人的数量
func (u *User) FollowingCount() int {
	return len(u.FollowingIDs())
}

// 关注的人的文章
func (u *User) FollowingPosts() (*[]Post, error) {
	var posts []Post
	ids := u.FollowingIDs()
	if err := db.Preload("User").
		Order("timestamp desc").
		Where("user_id in (?)", ids).
		Find(&posts).Error; err != nil {
		return nil, err
	}
	return &posts, nil
}

// 关注的人的文章,按照分页参数进行分页
func (u *User) FollowingPostsByPageAndLimit(page, limit int) (*[]Post, int, error) {
	var total int
	var posts []Post
	offset := (page - 1) * limit
	ids := u.FollowingIDs()
	if err := db.Preload("User").Order("timestamp desc").
		Where("user_id in (?)", ids).
		Offset(offset).
		Limit(limit).
		Find(&posts).Error; err != nil {
		return nil, total, err
	}
	db.Model(&Post{}).Where("user_id in (?)", ids).Count(&total)
	return &posts, total, nil
}

// 是否被某个用户关注
func (u *User) IsFollowedByUser(username string) bool {
	var count int
	user, _ := GetUserByUsername(username)
	if err := db.Table("follower").
		Where("user_id = ? and follower_id = ?", user.ID, u.ID).
		Count(&count).
		Error; err != nil {
		return false
	}
	return count > 0
}

// 创建文章
func (u *User) CreatePost(body string) error {
	post := Post{Body: body, UserID: u.ID}
	return db.Create(&post).Error
}

// 创建关于该用户的JWT字符串
func (u *User) GenerateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": u.Username,
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte("secret"))
}

// 检查JWT是否合法
func CheckToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("JWT加密算法错误,%v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["username"].(string), nil
	} else {
		return "", err
	}
}
