package main

import (
	"fmt"
	"reflect"
)

func reflectTypeAndValue(x interface{}) {
	tp := reflect.TypeOf(x)
	fmt.Printf("type:%v\n", tp)
	val := reflect.ValueOf(x)
	fmt.Println("value:", val)
}

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	ID   int    `json:"id"`
}

func main() {
	reflectTypeAndValue(10)
	reflectTypeAndValue(3.14)
	var s1 = Student{
		Name: "tom",
		Age:  18,
		ID:   100,
	}
	reflectTypeAndValue(s1)
}
