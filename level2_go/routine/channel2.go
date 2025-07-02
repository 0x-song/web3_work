package main

import (
	"fmt"
)

func main() {
	var ch1 = make(chan int, 5)
	for i := 0; i < 5; i++ {
		ch1 <- i + 1
	}
	close(ch1)

	for val := range ch1 {
		fmt.Println(val)
	}
}
