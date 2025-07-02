package main

import (
	"fmt"
	"time"
)

func main(){
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func(){
		for {
			ch1 <- 1
			time.Sleep(time.Second * 3)
		}
	}()

	go func(){
		for {
			ch2 <- 2
			time.Sleep(time.Second * 5)
		}
	}()

	for {
		select {
		case val := <-ch1:
			fmt.Println("接收到ch1的数据：", val)
		case val := <-ch2:
			fmt.Println("接收到ch2的数据：", val)
		}
	}

}