package main

import "fmt"

type Usber interface {
	Start()
	Stop()
}
type Phone struct {
	Name string
}

func (p Phone) Start() {
	fmt.Println(p.Name, "手机开始工作")
}
func (p Phone) Stop() {
	fmt.Println(p.Name, "手机停止工作")
}

type Camera struct {
}

func (c Camera) Start() {
	fmt.Println("相机开始工作")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作")
}

type Computer struct {
	Name string
}

func (c Computer) Working(usb Usber) {
	usb.Start()
	usb.Stop()
}

func main() {
	phone := Phone{
		Name: "小米",
	}
	var p Usber = phone
	p.Start()
	p.Stop()

	camera := Camera{}
	var c Usber = camera
	c.Start()
	c.Stop()

	fmt.Println("===================")
	computer := Computer{
		Name: "电脑",
	}
	computer.Working(phone)
	computer.Working(camera)
}
