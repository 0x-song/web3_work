package main

import "fmt"

func main() {
	s1 := []int{100, 200, 300}
	s2 := []int{400, 500, 600}
	s3 := append(s1, s2...)
	fmt.Println(s3)
}
