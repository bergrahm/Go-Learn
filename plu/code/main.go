package main

import (
	"c/Users/User/Desktop/repos/go/code/controllers"
	"net/http"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
