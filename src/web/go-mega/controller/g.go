package controller

import (
	"html/template"
	"web/go-mega/config"

	"github.com/gorilla/sessions"
)

// Author:Boyn
// Date:2020/2/23

var (
	homeController home                          // 定义homeController作为控制器
	templates      map[string]*template.Template // templates是页面控制器,通过扫描templates文件夹下的文件,进行一一映射
	sessionName    string                        // session变量的名字
	store          *sessions.CookieStore         // session存储的位置
)

func init() {
	templates = populateTemplates()
	store = sessions.NewCookieStore([]byte(config.GetSessionKey()))
	sessionName = "go-mega"
}

// Startup func
func Startup() {
	homeController.registerRoutes()
}
