package main

import (
	"fmt"
	"web/go-mega/model"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Author:Boyn
// Date:2020/2/24

func main() {
	fmt.Println("DB Init...")
	db := model.ConnectToDB()
	defer db.Close()
	model.SetDB(db)

	db.DropTableIfExists(model.User{}, model.Post{})
	db.CreateTable(model.User{}, model.Post{})
}
