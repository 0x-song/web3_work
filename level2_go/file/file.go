package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("1.txt")
	defer file.Close()
	if err != nil {
		fmt.Println("打开文件错误", err)
		return
	}
	fmt.Println(file)
}
