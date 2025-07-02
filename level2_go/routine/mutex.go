package main

import (
	"fmt"
	//"sync"
	"time"
)

var count = 0

func test() {
	count++
	fmt.Println("count = ", count)
	time.Sleep(time.Second)
}

func main() {
	for i := 0; i < 100; i++ {
		go test()
	}
	time.Sleep(time.Second * 5)
}
