package main

import (
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

	model.AddUser("bonfy", "abc123", "i@bonfy.im")
	model.AddUser("rene", "abc123", "rene@test.com")
	model.AddUser("boyn", "123456", "1065547951@qq.com")

	u1, _ := model.GetUserByUsername("bonfy")
	u1.CreatePost("Beautiful day in Portland!")
	model.UpdateAboutMe(u1.Username, `I'm the author of Go-Mega Tutorial you are reading now!`)

	u2, _ := model.GetUserByUsername("rene")
	u2.CreatePost("The Avengers movie was so cool!")
	u2.CreatePost("Sun shine is beautiful")

	u3, _ := model.GetUserByUsername("boyn")
	u3.CreatePost("Go Mega-Web开发编程实战")
	u3.CreatePost("Go语言")

	u1.Follow(u2.Username)
	u1.Follow(u3.Username)
	u2.Follow(u3.Username)

}
