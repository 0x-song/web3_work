package main

import (
	"fmt"
	"strconv"
	"time"
)

func test() {
	for i := 0; i < 30; i++ {
		fmt.Println("test() hello world " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {
	go test() // 开启一个协程，去执行test()
	for i := 0; i < 10; i++ {
		fmt.Println("main() hello golang " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
