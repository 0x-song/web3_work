package main

import "fmt"

func main() {
	var a = 10
	var b = &a
	fmt.Printf("a:%d &a:%p\n", a, &a)  // a:10 ptr:0xc0000100a8
	fmt.Printf("b:%p type:%T\n", b, b) // b:0xc0000100a8 type:*int
	fmt.Println("取b的地址:", &b)          // 0xc000006028
}
