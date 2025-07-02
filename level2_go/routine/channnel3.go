package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		fmt.Println("准备接收数据")
		<-ch
		fmt.Println("接收到数据")
	}()

	time.Sleep(time.Second * 40)
}
