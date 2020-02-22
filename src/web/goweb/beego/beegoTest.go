package main

import (
	"database/sql"
	"fmt"

	"github.com/astaxie/beedb"
	_ "github.com/go-sql-driver/mysql"
)

// Author:Boyn
// Date:2020/2/22

func main() {
	db, err := sql.Open("mysql", "waytogo:123456@/way_to_go?charset=utf8mb4")
	if err != nil {
		panic(err)
	}
	orm := beedb.New(db)
	/*var saveone UserInfo
	saveone.Username = "Boyn"
	saveone.Departname = "ABCD"
	saveone.Created = time.Now()

	// 使用单个结构体进行添加
	err = orm.Save(&saveone)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(saveone.Uid)


	addSlice := make([]map[string]interface{},0)
	add := make(map[string]interface{})
	add2 := make(map[string]interface{})
	add["username"] = "AAA"
	add["departname"] = "某部门"
	add["created"] = time.Now()
	add2["username"] = "AAA"
	add2["departname"] = "某部门"
	add2["created"] = time.Now()
	addSlice = append(addSlice,add,add2)
	//使用Map进行批量添加
	batch, err := orm.SetTable("user_info").InsertBatch(addSlice)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(batch)*/

	var user UserInfo
	orm.Where("uid=?", 9).Find(&user)
	fmt.Println(user)

	users := make([]UserInfo, 0)
	err = orm.Where("username = ?", "boyn").FindAll(&users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(users)

	users = make([]UserInfo, 0)
	err = orm.Where("username = ? and departname = ?", "boyn", "01").FindAll(&users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(users)
}
