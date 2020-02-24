package controller

import (
	"log"
	"net/http"
	"web/go-mega/model"
)

// Author:Boyn
// Date:2020/2/24

// 中间层,用于判断用户是否登录
// 如果未登录,则将其重定向至登录页面中
// 如果已经登录了,就更新用户的last_seen字段,并继续下一步处理
func middleAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		username, err := getSessionUser(request)
		log.Println("middle:", username)
		if err != nil {
			log.Println("中间层无法获取到缓存")
			http.Redirect(writer, request, "/login", 302)
		} else {
			// 更新用户的last seen时间
			model.UpdateLastSeen(username)
			next.ServeHTTP(writer, request)
		}
	}
}
