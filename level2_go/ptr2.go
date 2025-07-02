package main

import "fmt"

func main() {
	// var userinfo map[string]string
	// userinfo["username"] = "张三"
	// fmt.Println(userinfo)

	var a *int
	*a = 100
	fmt.Println(*a)
}
