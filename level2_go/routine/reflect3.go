package main

import (
	"fmt"
	"reflect"
)

func reflectSet(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(20)
	}
}

func reflectSet2(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(20)
	}
}

func main() {
	var a int64 = 10
	fmt.Println(a)
	reflectSet2(&a)
	fmt.Println(a)
}
