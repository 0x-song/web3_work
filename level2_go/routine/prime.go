package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func putNum(intChan chan int) {
	for i := 1; i <= 12000; i++ {
		intChan <- i
	}
	close(intChan)
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool, id int) {

}

func main() {
	start := time.Now().Unix()
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 20000)
	exitChan := make(chan bool, 8)
	//将12000个数字放入到一个channel中
	putNum(intChan)

	//开启4个协程，从intChan取出数据，判断是否为素数，如果是就放入到primeChan中
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go primeNum(intChan, primeChan, exitChan, i)
	}

	end := time.Now().Unix()
	fmt.Println("执行完毕,耗时：", end-start, "秒")
}
