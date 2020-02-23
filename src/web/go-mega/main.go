package main

import (
	"html/template"
	"net/http"
)

// Author:Boyn
// Date:2020/2/23
type User struct {
	Username string
}

type Post struct {
	User  User
	Title string
}

type IndexViewModel struct {
	User  User
	Title string
	Posts []Post
}

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		u1 := User{Username: "boyn"}
		u2 := User{Username: "kaserchan"}
		posts := []Post{
			{
				User:  u1,
				Title: "模板的if-else操作",
			},
			{
				User:  u1,
				Title: "模板的插值操作",
			},
			{
				User:  u2,
				Title: "模板的循环操作",
			},
		}
		v := IndexViewModel{
			User:  u1,
			Title: "Home Page",
			Posts: posts,
		}
		tpl, _ := template.ParseFiles("F:\\Code\\Go\\LearningGo\\src\\web\\go-mega\\templates\\index.html")
		tpl.Execute(writer, &v)
	})
	http.ListenAndServe(":8009", nil)
}
