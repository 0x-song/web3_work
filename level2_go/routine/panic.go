package main

import (
	"fmt"
)

// 函数
func test() {
	//这里我们可以使用defer + recover
	defer func() {
		//捕获 test 抛出的 panic
		if err := recover(); err != nil {
			fmt.Println("test() 发生错误", err)
		}
	}()
	//定义了一个 map
	var myMap map[int]string
	myMap[0] = "golang" //error
}
func main() {
	test()
	fmt.Println("main()下面的代码")
}
