package main

import (
	"fmt"
	"sync"
	"time"
)

func f1() {
	start := time.Now().Unix()
	for num := 0; num < 120000; num++ {
		flag := true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			//fmt.Println(num)
		}
	}
	end := time.Now().Unix()

	fmt.Println("f1执行时间", end-start)
}

var wg sync.WaitGroup

func f2(n int) {
	for num := (n - 1) * 30000; num <= n*30000; num++ {
		flag := true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			//fmt.Println(num)
		}
	}
	wg.Done()
}

func main() {
	f1()
	start := time.Now().Unix()
	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go f2(i)
	}
	wg.Wait()
	end := time.Now().Unix()
	fmt.Println("f2执行时间", end-start)
}
