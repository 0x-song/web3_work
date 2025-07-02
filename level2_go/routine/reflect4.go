package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

func (s Student) GetInfo() string {
	var str = fmt.Sprintf("姓名:%v 年龄:%v 成绩:%v", s.Name, s.Age, s.Score)
	fmt.Println(str)
	return str
}

func (s *Student) SetInfo(name string, age int, score int) {
	s.Name = name
	s.Age = age
	s.Score = score
}

func printStructField(s interface{}) {
	t := reflect.TypeOf(s)
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("数据类型不正确")
		return
	}

	field0 := t.Field(0)
	fmt.Printf("field0: %v\n", field0)
	fmt.Println(field0.Name)
	fmt.Println(field0.Type)
	fmt.Println(field0.Tag.Get("json"))

	field1, _ := t.FieldByName("Age")
	fmt.Printf("field1: %v\n", field1)
	fmt.Println(field1.Name)
	fmt.Println(field1.Type)
	fmt.Println(field1.Tag.Get("json"))

	num := t.NumField()
	fmt.Printf("结构体 %v 的字段个数为：%v\n", t.Name(), num)
}

// 方法
func PrintStructFn(s interface{}) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("数据类型不正确")
		return
	}

	num := t.NumMethod()
	fmt.Printf("结构体 %v 的方法个数为：%v\n", t.Name(), num)

	method := t.Method(0)
	fmt.Printf("方法名：%v\n", method.Name)
	fmt.Printf("方法类型：%v\n", method.Type)
	// 调用方法
	//var args = []reflect.Value{}
	result := v.Method(0).Call(nil) //GetInfo
	fmt.Println("GetInfo调用结果:", result)
	// 调用方法

	var params []reflect.Value
	params = append(params, reflect.ValueOf("张三"))
	params = append(params, reflect.ValueOf(18))
	params = append(params, reflect.ValueOf(100))
	v.MethodByName("SetInfo").Call(params)

}

func main() {
	var s1 = Student{
		Name:  "tom",
		Age:   18,
		Score: 90,
	}
	//printStructField(s1)
	PrintStructFn(&s1)
	fmt.Println(s1)
}
