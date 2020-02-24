package main

import (
	"fmt"
	"net/http"
	"web/go-mega/controller"
	"web/go-mega/model"

	"github.com/gorilla/context"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Author:Boyn
// Date:2020/2/23

func main() {
	// Set up DB
	fmt.Println("DB Init...")
	db := model.ConnectToDB()
	defer db.Close()
	model.SetDB(db)

	// Set up Controller
	controller.Startup()
	http.ListenAndServe(":8009", context.ClearHandler(http.DefaultServeMux))
}
