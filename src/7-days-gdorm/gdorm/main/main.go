package main

import (
	gdorm "7-days-gdorm/gdorm"
	_ "github.com/mattn/go-sqlite3"
)

// Author:Boyn
// Date:2020/3/23

func main() {
	engine, _ := gdorm.NewEngine("sqlite3", "/Users/chenzhanpeng/Code/Go/LearningGo/src/7-days-gdorm/gdorm/gee.db")
	defer engine.Close()
	s := engine.NewSession()
	_, _ = s.Raw("drop table if exists User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text );").Exec()
	_, _ = s.Raw("insert into User values(?)", "Jack").Exec()
	_, _ = s.Raw("insert into User values(?)", "Bob").Exec()
}
