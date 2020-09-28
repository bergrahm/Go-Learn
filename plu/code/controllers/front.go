package controllers

import "net/http"

func RegisterControllers() {
	uc := newUserController()

	http.Handle("/Users", *uc)
	http.Handle("/Users/", *uc)
}
