package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func test() {
	for i := 0; i < 30; i++ {
		fmt.Println("test() hello world " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
	wg.Done()
}
func main() {
	wg.Add(1) //表示有一个协程需要等待
	go test() // 开启一个协程，去执行test()
	for i := 0; i < 10; i++ {
		fmt.Println("main() hello golang " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
	wg.Wait()
}
