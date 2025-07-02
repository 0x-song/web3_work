package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	ID     int
	Gender string
	name   string
	Sno    string
}

func main() {
	var s1 = Student{
		ID:     1,
		Gender: "男",
		name:   "小王子",
		Sno:    "101010",
	}
	fmt.Printf("%#v\n", s1)
	var s, _ = json.Marshal(s1)
	fmt.Println(string(s))

	//var jsonStr = "{\"ID\":1,\"Gender\":\"男\",\"name\":\"小王子\",\"Sno\":\"101010\"}"
	var jsonStr2 = `{"ID":1,"Gender":"男","name":"小王子","Sno":"101010"}`
	var student Student
	err := json.Unmarshal([]byte(jsonStr2), &student)
	if err != nil {
		fmt.Printf("unmarshal failed, err:%v\n", err)
	} else {
		fmt.Printf("%#v\n", student)
	}
}
