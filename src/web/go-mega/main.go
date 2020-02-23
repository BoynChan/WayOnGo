package main

import (
	"net/http"
	"web/go-mega/controller"
)

// Author:Boyn
// Date:2020/2/23

func main() {
	controller.Startup()
	http.ListenAndServe(":8009", nil)
}
