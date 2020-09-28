package main

import (
	"c/Users/User/Desktop/repos/go/code/models"
	"fmt"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Tricia",
		LastName:  "McMilla",
	}
	fmt.Println(u)
}
