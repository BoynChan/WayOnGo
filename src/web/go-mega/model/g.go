package model

import (
	"fmt"
	"web/go-mega/config"

	"github.com/jinzhu/gorm"
)

// Author:Boyn
// Date:2020/2/24

var db *gorm.DB

func SetDB(database *gorm.DB) {
	db = database
}

func ConnectToDB() *gorm.DB {
	connectionString := config.GetMysqlConnectionString()
	fmt.Println("Connect to db ... ")
	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		panic(fmt.Sprintf("无法连接数据库,%s", err))
	}
	db.SingularTable(true)
	return db
}
