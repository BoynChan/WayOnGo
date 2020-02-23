package controller

import (
	"net/http"
	"web/go-mega/vm"
)

// Author:Boyn
// Date:2020/2/23
type home struct{}

func (h home) registerRoutes() {
	http.HandleFunc("/", indexHandler)
}

// 这个函数被registerRoutes()注册在homeController中,它注册的路径是'/'
// 当访问到这个路径的时候,会使用这个函数来进行View返回
func indexHandler(w http.ResponseWriter, r *http.Request) {
	vop := vm.IndexViewModelOp{}
	v := vop.GetVM()
	templates["index.html"].Execute(w, &v)
}
