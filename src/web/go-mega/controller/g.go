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

// controller下的初始化函数
func init() {
	// 初始化要用到的网页模板
	templates = populateTemplates()

	// 初始化session工具
	store = sessions.NewCookieStore([]byte(config.GetSessionKey()))

	//设置session的名字为go-mega,这个名字同时也会反映到cookie中
	sessionName = "go-mega"
}

// Startup func
func Startup() {
	// 将各个controller中声明的路由注册到http中
	homeController.registerRoutes()
}
