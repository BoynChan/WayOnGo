package main

import "net/http"

// Author:Boyn
// Date:2020/2/23

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello!"))
	})
	http.ListenAndServe(":8009", nil)
}
