package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Author:Boyn
// Date:2020/2/21

func main() {
	db, err := sql.Open("mysql", "waytogo:123456@/way_to_go?charset=utf8mb4")
	checkErr(err)
	/*stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	checkErr(err)
	res, err := stmt.Exec("boyn", "研发部门", "2012-12-09")
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)*/

	rows, err := db.Query("SELECT * FROM userinfo")
	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid, username, department, created)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
