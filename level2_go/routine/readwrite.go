package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 写数据
func fn1(ch chan int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch <- i + 1
		fmt.Println("写入数据：", i+1)
		time.Sleep(time.Second)
	}
	close(ch)
}

func fn2(ch chan int) {
	defer wg.Done()
	fmt.Println("准备接收数据")
	for v := range ch {
		fmt.Println("读取数据：", v)
		time.Sleep(time.Millisecond * 500)
	}
}

// 读数据
func main() {
	ch := make(chan int)
	wg.Add(2)
	go fn1(ch)
	go fn2(ch)
	wg.Wait()
	fmt.Println("读取完毕")
}
