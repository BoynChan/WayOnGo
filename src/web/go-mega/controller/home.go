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
	http.HandleFunc("/login", loginHandler)
}

// 这个函数被registerRoutes()注册在homeController中,它注册的路径是'/'
// 当访问到这个路径的时候,会使用这个函数来进行View返回
func indexHandler(w http.ResponseWriter, r *http.Request) {
	vop := vm.IndexViewModelOp{}
	v := vop.GetVM()
	templates["index.html"].Execute(w, &v)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	temName := "login.html"
	vop := vm.LoginViewModelOp{}
	v := vop.GetVM()
	if r.Method == http.MethodGet {
		templates[temName].Execute(w, &v)
	} else if r.Method == http.MethodPost {
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		if len(username) < 3 {
			v.AddError("用户名过短")
		}
		if len(password) < 6 {
			v.AddError("密码过短")
		}
		if !check(username, password) {
			v.AddError("用户名与密码错误")
		}
		if len(v.Errs) > 0 {
			templates[temName].Execute(w, &v)
		} else {
			//如果密码正确,就进行302重定向
			http.Redirect(w, r, "/", 302)
		}
	}
}

func check(username, password string) bool {
	if username == "boyn" && password == "123456" {
		return true
	}
	return false
}
