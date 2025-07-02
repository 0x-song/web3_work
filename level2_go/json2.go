package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	ID     int    `json:"id"`
	Gender string `json:"gender"`
	Name   string `json:"name"`
	Sno    string
}

type Person struct {
	ID     int
	Gender string
	Name   string
	Sno    string
}

func main() {
	s1 := Student{
		ID:     1,
		Gender: "男",
		Name:   "张三",
		Sno:    "123",
	}
	fmt.Println(s1)
	var s, _ = json.Marshal(s1)
	jsonStr := string(s)
	fmt.Println(jsonStr)

	var p = Person{
		ID:     1,
		Gender: "男",
		Name:   "李四",
		Sno:    "123",
	}
	var s2, _ = json.Marshal(p)
	jsonStr2 := string(s2)
	fmt.Println(jsonStr2)
}
