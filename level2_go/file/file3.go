package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("1.txt")
	if err != nil {
		fmt.Println("open file err=", err)
	}
	defer file.Close()

	// 创建一个 *Reader, 是带缓冲的
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			fmt.Println("文件读完了")
			break
		}
		if err!= nil {
			fmt.Println("read file err=", err)
			break
		}
		fmt.Println(line)
	}

}
