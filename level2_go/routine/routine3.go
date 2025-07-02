package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done() //结束就登记-1
	fmt.Println("hello", i)
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1) //启动一个协程就登记+1
		go hello(i)
	}

	wg.Wait()
}
