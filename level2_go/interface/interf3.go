package main

import "fmt"

func main() {
	var x interface{}
	x = "Hello golnag"
	v, ok := x.(string)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("类型断言失败")
	}
}
