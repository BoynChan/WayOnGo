package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Author:Boyn
// Date:2020/3/23
// 简单的Sqlite操作实例

func main() {
	db, _ := sql.Open("sqlite3", "../gee.db")
	defer db.Close()
	_, _ = db.Exec("drop table if exists User;")
	_, _ = db.Exec("CREATE TABLE User(Name text,Age integer );")
	result, err := db.Exec("insert into User values(?,?)", "Tom", 18)
	if err == nil {
		affected, _ := result.RowsAffected()
		fmt.Println(affected)
	}
	result, err = db.Exec("insert into User values(?,?)", "Jack", 28)
	if err == nil {
		affected, _ := result.RowsAffected()
		fmt.Println(affected)
	}
	row := db.QueryRow("select Name from User limit 1")
	var name string
	if err := row.Scan(&name); err == nil {
		fmt.Println(name)
	}
}
