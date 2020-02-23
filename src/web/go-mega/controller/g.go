package controller

import "html/template"

// Author:Boyn
// Date:2020/2/23

var (
	homeController home                          // 定义homeController作为控制器
	templates      map[string]*template.Template // templates是页面控制器,通过扫描templates文件夹下的文件,进行一一映射
)

func init() {
	templates = populateTemplates()
}

// Startup func
func Startup() {
	homeController.registerRoutes()
}
