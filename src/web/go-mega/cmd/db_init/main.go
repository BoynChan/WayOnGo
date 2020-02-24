package main

import (
	"fmt"
	"log"
	"web/go-mega/model"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Author:Boyn
// Date:2020/2/24

func main() {
	log.Println("DB Init ...")
	db := model.ConnectToDB()
	defer db.Close()
	model.SetDB(db)

	db.DropTableIfExists(model.User{}, model.Post{})
	db.CreateTable(model.User{}, model.Post{})

	users := []model.User{
		{
			Username:     "bonfy",
			PasswordHash: model.GeneratePasswordHash("abc123"),
			Email:        "i@bonfy.im",
			Avatar:       fmt.Sprintf("https://www.gravatar.com/avatar/%s?d=identicon", model.Md5("i@bonfy.im")),
			Posts: []model.Post{
				{Body: "Beautiful day in Portland!"},
			},
		},
		{
			Username:     "rene",
			PasswordHash: model.GeneratePasswordHash("abc123"),
			Email:        "rene@test.com",
			Avatar:       fmt.Sprintf("https://www.gravatar.com/avatar/%s?d=identicon", model.Md5("rene@test.com")),
			Posts: []model.Post{
				{Body: "The Avengers movie was so cool!"},
				{Body: "Sun shine is beautiful"},
			},
		},
		{
			Username:     "boyn",
			PasswordHash: model.GeneratePasswordHash("123456"),
			Email:        "1065547951@qq.com",
			Avatar:       fmt.Sprintf("https://www.gravatar.com/avatar/%s?d=identicon", model.Md5("1065547951@qq.com")),
			Posts: []model.Post{
				{Body: "Go语言"},
				{Body: "Web开发"},
			},
		},
	}

	for _, u := range users {
		db.Debug().Create(&u)
	}
}
