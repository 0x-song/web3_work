package main

import (
	"fmt"
	"reflect"
)

func reflectType(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("TypeOf:%v Name:%v Kind:%v\n", t, t.Name(), t.Kind())
}

type myInt int64
type Person struct {
	Name string
	Age  int
}
type Animal struct {
	Name string
}

func main() {
	var a *float32 // 指针
	var b myInt    // 自定义类型
	var c rune     // 类型别名
	reflectType(a) // type: kind:ptr
	reflectType(b) // type:myInt kind:int64
	reflectType(c) // type:int32 kind:int32
	var d = Person{
		Name: "itying",
		Age:  18,
	}
	var e = Animal{Name: "小花"}
	reflectType(d) // type:Person kind:struct
	reflectType(e) // type:Animal kind:struct
	var f = []int{1, 2, 3, 4, 5}
	reflectType(f) //TypeOf:[]int Name: Kind:slice
}
