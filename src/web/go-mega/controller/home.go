package controller

import (
	"fmt"
	"log"
	"net/http"
	"web/go-mega/vm"

	"github.com/gorilla/mux"
)

// Author:Boyn
// Date:2020/2/23
type home struct{}

// 注册 homeController 的路由
// 在这里将其设置为home的方法,是因为通常有多个controller
// 以示区分
func (h home) registerRoutes() {
	r := mux.NewRouter()

	r.HandleFunc("/", middleAuth(indexHandler))
	r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/register", registerHandler)
	r.HandleFunc("/logout", middleAuth(logoutHandler))
	r.HandleFunc("/user/{username}", middleAuth(profileHandler))

	http.Handle("/", r)
}

// 这个函数被registerRoutes()注册在homeController中,它注册的路径是'/'
// 当访问到这个路径的时候,会使用这个函数来进行View返回
// 根据已经登录的用户来获取它的用户名
func indexHandler(w http.ResponseWriter, r *http.Request) {
	vop := vm.IndexViewModelOp{}
	username, _ := getSessionUser(r)
	v := vop.GetVM(username)
	templates["index.html"].Execute(w, &v)
}

// 登录的处理函数
// 如果是get请求,就返回登录的网页
// 如果是post请求,就验证密码是否正确
// 如果不正确,将错误信息发送到模板中,并返回登录模板网页
// 如果正确,就设置缓存,并进行302重定向到首页
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

		errs := checkLogin(username, password)
		v.AddError(errs...)
		if len(v.Errs) > 0 {
			templates[temName].Execute(w, &v)
		} else {
			setSessionUser(w, r, username)
			//如果密码正确,就进行302重定向
			http.Redirect(w, r, "/", 302)
		}
	}
}

// 登出的处理函数
// 登出时,先将缓存清除,然后重定向到首页
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	clearSession(w, r)
	http.Redirect(w, r, "/", 302)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	temName := "register.html"
	vop := vm.RegisterModelViewOp{}
	v := vop.GetVM()
	if r.Method == http.MethodGet {
		templates[temName].Execute(w, &v)
		return
	}
	if r.Method == http.MethodPost {
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		email := r.Form.Get("email")
		errs := checkRegister(username, email, password)
		v.AddError(errs...)

		if len(v.Errs) > 0 {
			templates[temName].Execute(w, &v)
		} else {
			if err := addUser(username, password, email); err != nil {
				log.Println("add User error:", err)
				w.Write([]byte("插入数据库错误"))
				return
			}
			setSessionUser(w, r, username)
			http.Redirect(w, r, "/", 302)
		}
	}
}

// 关于用户个人资料的controller
// 这里用到了动态路由的库
// mux.Vars获取路由信息中被动态代理的消息,并通过map取出
// pUser表示要查看个人档案的用户,sUser表示当前的用户
func profileHandler(w http.ResponseWriter, r *http.Request) {
	temName := "profile.html"
	vars := mux.Vars(r)
	pUser := vars["username"]
	sUser, _ := getSessionUser(r)
	vop := vm.ProfileViewModelOp{}
	v, err := vop.GetVM(sUser, pUser)
	if err != nil {
		msg := fmt.Sprintf("用户 %s 不存在", pUser)
		w.Write([]byte(msg))
		return
	}
	templates[temName].Execute(w, &v)
}
