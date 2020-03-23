package main

import (
	gdorm "7-days-gdorm/day1-database-sql"
	"7-days-gdorm/day1-database-sql/log"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Author:Boyn
// Date:2020/3/23
type User struct {
	Name string
	Age  int
}

func main() {
	engine, _ := gdorm.NewEngine("sqlite3", "../gee.db")
	defer engine.Close()
	s := engine.NewSession()
	_, _ = s.Raw("drop table if exists User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text );").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	_, _ = s.Raw("insert into User values(?)", "Tom").Exec()
	_, _ = s.Raw("insert into User values(?)", "Java").Exec()
	row := s.Raw("select Name from User limit 1").QueryRow()
	var name string
	if err := row.Scan(&name); err == nil {
		fmt.Println(name)
	} else {
		log.Error(err)
	}
}
