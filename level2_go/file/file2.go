package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("1.txt")
	if err != nil {
		fmt.Println("open file err=", err)
		return
	}
	defer file.Close()

	var buf = make([]byte, 1024)
	num, err2 := file.Read(buf)
	if err == io.EOF {
		fmt.Println("文件读取完毕")
		return
	}
	if err2 != nil {
		fmt.Println("read file err=", err2)
		return
	}
	fmt.Printf("读取到了%d个字节\n", num)
	fmt.Printf("读取到的内容为：%v\n", string(buf[:num]))
}
