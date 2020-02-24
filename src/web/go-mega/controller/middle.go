package controller

import (
	"log"
	"net/http"
)

// Author:Boyn
// Date:2020/2/24

// 中间层,用于判断用户是否登录
func middleAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		username, err := getSessionUser(request)
		log.Println("middle:", username)
		if err != nil {
			log.Println("中间层无法获取到缓存")
			http.Redirect(writer, request, "/login", 302)
		} else {
			next.ServeHTTP(writer, request)
		}
	}
}
