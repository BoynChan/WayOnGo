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

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		user := User{Username: "boyn"}
		tpl, _ := template.New("").Parse(`
		<html><head><title>Home - Boyn</title></head>
<body>
<h1>Hello,{{.Username}}</h1>
</body>
`)
		tpl.Execute(writer, &user)

	})
	http.ListenAndServe(":8009", nil)
}
