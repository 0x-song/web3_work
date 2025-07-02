package main

import (
	"fmt"
	"sync"
	"time"
)

var count = 0
var mutex sync.Mutex

func test() {

	mutex.Lock()
	defer mutex.Unlock()

	count++
	fmt.Println("count = ", count)
	time.Sleep(time.Millisecond)
}

func main() {
	for i := 0; i < 100; i++ {
		go test()
	}
	time.Sleep(time.Second * 5)
}
