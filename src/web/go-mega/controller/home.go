package controller

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
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
	r.HandleFunc("/profile_edit", middleAuth(profileEditHandler))
	r.HandleFunc("/follow/{username}", middleAuth(followHandler))
	r.HandleFunc("/unfollow/{username}", middleAuth(unFollowHandler))
	r.HandleFunc("/reset_password_request", resetPasswordRequestHandler)
	r.HandleFunc("/explore", exploreHandler)
	r.HandleFunc("/reset_password/{token}", resetPasswordHandler)

	http.Handle("/", r)
}

// 主页的handler
// 主页目前的功能是可以显示最近发表的文章和发表新的文章
// 当请求为GET时,返回规划的VM
// 当请求为POST时,用户发布动态,将其插入到数据库中
func indexHandler(w http.ResponseWriter, r *http.Request) {
	temName := "index.html"
	vop := vm.IndexViewModelOp{}
	page := getPage(r)
	username, _ := getSessionUser(r)
	if r.Method == http.MethodGet {
		flash := getFlash(w, r)
		v := vop.GetVM(username, flash, page, pageLimit)
		_ = templates[temName].Execute(w, &v)
	}
	if r.Method == http.MethodPost {
		_ = r.ParseForm()
		body := r.Form.Get("body")
		errMessage := checkLen("Post", body, 1, 180)
		if len(errMessage) != 0 {
			setFlash(w, r, strings.Join(errMessage, "\n"))
		} else {
			err := vm.CreatePost(username, body)
			if err != nil {
				log.Println("add Post error:", err)
				_, _ = w.Write([]byte("插入新文章失败"))
				return
			}
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
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
		_ = templates[temName].Execute(w, &v)
	} else if r.Method == http.MethodPost {
		_ = r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")

		errs := checkLogin(username, password)
		v.AddError(errs...)
		if len(v.Errs) > 0 {
			_ = templates[temName].Execute(w, &v)
		} else {
			// 设置用户为登录状态
			_ = setSessionUser(w, r, username)
			//如果密码正确,就进行302重定向
			http.Redirect(w, r, "/", 302)
		}
	}
}

// 登出的处理函数
// 登出时,先将缓存清除,然后重定向到首页
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	_ = clearSession(w, r)
	http.Redirect(w, r, "/", 302)
}

// 注册的handler
// 当为GET方法时,即将返回注册页面
// 当为POST方法时,解析参数并检查
// 如果发生检查错误,则将错误信息打在页面上
// 如果没有错误,则重定向到首页,并定义用户为登录状态
func registerHandler(w http.ResponseWriter, r *http.Request) {
	temName := "register.html"
	vop := vm.RegisterModelViewOp{}
	v := vop.GetVM()
	if r.Method == http.MethodGet {
		_ = templates[temName].Execute(w, &v)
		return
	}
	if r.Method == http.MethodPost {
		_ = r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		email := r.Form.Get("email")
		errs := checkRegister(username, email, password)
		v.AddError(errs...)

		if len(v.Errs) > 0 {
			_ = templates[temName].Execute(w, &v)
		} else {
			if err := addUser(username, password, email); err != nil {
				log.Println("add User error:", err)
				_, _ = w.Write([]byte("插入数据库错误"))
				return
			}
			_ = setSessionUser(w, r, username)
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
	page := getPage(r)
	pUser := vars["username"]
	sUser, _ := getSessionUser(r)
	vop := vm.ProfileViewModelOp{}
	v, err := vop.GetVM(sUser, pUser, page, pageLimit)
	if err != nil {
		msg := fmt.Sprintf("用户 %s 不存在", pUser)
		_, _ = w.Write([]byte(msg))
		return
	}
	_ = templates[temName].Execute(w, &v)
}

// 编辑个人资料的handler
// get方法会返回个人资料VM
// post会更新个人资料并重定向至个人页面
func profileEditHandler(w http.ResponseWriter, r *http.Request) {
	temName := "profile_edit.html"
	username, _ := getSessionUser(r)
	vop := vm.ProfileEditViewModelOP{}
	v := vop.GetVM(username)
	if r.Method == http.MethodGet {
		_ = templates[temName].Execute(w, &v)
	}
	if r.Method == http.MethodPost {
		_ = r.ParseForm()
		aboutme := r.Form.Get("aboutme")
		log.Println("[ProfileEditHandler] about me :", aboutme)
		if err := vm.UpdateAboutMe(username, aboutme); err != nil {
			log.Println("[ProfileEditHandler] err :", err)
			_, _ = w.Write([]byte("更新时出现错误"))
			return
		}
		http.Redirect(w, r, fmt.Sprintf("/user/%s", username), 302)
	}
}

// 添加用户到关注列表功能
func followHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pUser := vars["username"]
	sUser, _ := getSessionUser(r)
	err := vm.Follow(sUser, pUser)
	if err != nil {
		log.Println("[followHandler] 关注失败:", err)
		_, _ = w.Write([]byte("关注失败"))
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/user/%s", pUser), http.StatusSeeOther)
}

// 取消关注功能
func unFollowHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pUser := vars["username"]
	sUser, _ := getSessionUser(r)
	err := vm.UnFollow(sUser, pUser)
	if err != nil {
		log.Println("[unFollowHandler] 取消关注失败:", err)
		_, _ = w.Write([]byte("取消关注失败"))
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/user/%s", pUser), http.StatusSeeOther)
}

// 探索功能的handler
// 类似于微博的广场功能,会显示所有用户的发送消息
func exploreHandler(w http.ResponseWriter, r *http.Request) {
	temName := "explore.html"
	vop := vm.ExploreViewModelOp{}
	page := getPage(r)
	v, err := vop.GetVM(page, pageLimit)
	if err != nil {
		log.Println("[exploreHandler] 获取广场文章失败:", err)
		_, _ = w.Write([]byte("获取广场文章失败"))
		return
	}
	_ = templates[temName].Execute(w, &v)
}

// 密码重置时,用于发送邮件功能的handler
func resetPasswordRequestHandler(w http.ResponseWriter, r *http.Request) {
	temName := "reset_password_request.html"
	vop := vm.ResetPasswordRequestViewModelOp{}
	v := vop.GetVM()

	if r.Method == http.MethodPost {
		r.ParseForm()
		email := r.Form.Get("email")

		errs := checkResetPasswordRequest(email)
		v.AddError(errs...)

		if len(v.Errs) > 0 {
			templates[temName].Execute(w, &v)
		} else {
			log.Println("[resetPasswordRequestHandler] 发送邮件至", email)
			vopEmail := vm.EmailViewModelOp{}
			vEmail := vopEmail.GetVM(email)
			var contentByte bytes.Buffer
			tem, _ := template.ParseFiles(".\\src\\web\\go-mega\\templates\\content\\email.html")
			if err := tem.Execute(&contentByte, &vEmail); err != nil {
				log.Println("[resetPasswordRequestHandler] 邮件发送失败:", err)
				w.Write([]byte("发送邮件错误"))
				return
			}
			content := contentByte.String()
			go sendEmail(email, "重置密码", content)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
	}

	if r.Method == http.MethodGet {
		templates[temName].Execute(w, &v)
	}
}

// 接收到邮件时,用于重置密码的功能handler
func resetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	token := vars["token"]
	username, err := vm.CheckToken(token)
	if err != nil {
		w.Write([]byte("The token is no longer valid, please go to the login page."))
	}

	tpName := "reset_password.html"
	vop := vm.ResetPasswordViewModelOp{}
	v := vop.GetVM(token)

	if r.Method == http.MethodGet {
		templates[tpName].Execute(w, &v)
	}

	if r.Method == http.MethodPost {
		log.Println("Reset password for ", username)
		r.ParseForm()
		pwd1 := r.Form.Get("pwd1")
		pwd2 := r.Form.Get("pwd2")

		errs := checkResetPassword(pwd1, pwd2)
		v.AddError(errs...)

		if len(v.Errs) > 0 {
			templates[tpName].Execute(w, &v)
		} else {
			if err := vm.ResetUserPassword(username, pwd1); err != nil {
				log.Println("reset User password error:", err)
				w.Write([]byte("Error update user password in database"))
				return
			}
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
	}
}
