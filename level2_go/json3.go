package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	ID     int
	Gender string
	Name   string
}
type Class struct {
	Title    string
	Students []Student
}

func main() {
	c := Class{
		Title:    "101",
		Students: make([]Student, 200),
	}

	for i := 0; i < 10; i++ {
		stu := Student{
			Name:   "stu",
			Gender: "男",
			ID:     i,
		}
		c.Students = append(c.Students, stu)
	}
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed")
	}
	fmt.Printf("json:%s\n", data)

	str := `{"Title":"101","Students":[{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"","Name":""},{"ID":0,"Gender":"男","Name":"stu"},{"ID":1,"Gender":"男","Name":"stu"},{"ID":2,"Gender":"男","Name":"stu"},{"ID":3,"Gender":"男","Name":"stu"},{"ID":4,"Gender":"男","Name":"stu"},{"ID":5,"Gender":"男","Name":"stu"},{"ID":6,"Gender":"男","Name":"stu"},{"ID":7,"Gender":"男","Name":"stu"},{"ID":8,"Gender":"男","Name":"stu"},{"ID":9,"Gender":"男","Name":"stu"}]}`

	c1 := &Class{}

	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("json unmarshal failed")
	}
	fmt.Printf("%#v\n", c1)

}
