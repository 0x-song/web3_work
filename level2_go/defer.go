package main

import "fmt"

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	//因为x是局部变量，返回值=x已经确定
	return x
}
func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 // return x 等价于 x=5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	//defer影响的是函数内部的x，不会对返回值有任何影响
	defer func(x int) {
		x++
	}(x)
	return 5
}

func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}
