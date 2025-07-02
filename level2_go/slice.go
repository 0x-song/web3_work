package main
import "fmt"

func main(){
	s := []int{2,3,5,7,11,13}
	fmt.Println(s)
	fmt.Printf("长度:%v 容量 %v\n", len(s), cap(s))
	c := s[:2]
	fmt.Println(c)
	fmt.Printf("长度:%v 容量 %v\n", len(c), cap(c))
	d := s[1:3]
	fmt.Println(d)
	fmt.Printf("长度:%v 容量 %v", len(d), cap(d))
}