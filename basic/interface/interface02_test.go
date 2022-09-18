package _interface

import (
	"fmt"
	"testing"
)

/*
面向对象编程思想-接口
「interface是引用类型」
1，在golang中多态特性主要是通过接口来实现的。
2，为什么要有接口？ 松耦合,高内聚的编程思想
*/

func TestInterface02(t *testing.T) {
	//接口快速入门
	interface21()
}

//声明/定义一个接口
type Usb interface {
	//声明了2个没有实现的方法
	Start()
	Stop()
}

type Phone struct {
	Name string
}

type Camera struct {
	Name string
}

type Computer struct {
}

//让phone实现usb接口的方法
func (p Phone) Start() {
	fmt.Println("phone start work.")
}

func (p Phone) Stop() {
	fmt.Println("phone stop work.")
}

//让Camera实现usb接口的方法
func (c Camera) Start() {
	fmt.Println("Camera start work.")
}

func (c Camera) Stop() {
	fmt.Println("Camera stop work.")
}

//此方法接收一个Usb接口类型变量，只要是实现了Usb接口.「所谓实现Usb接口，就是指实现了Usb接口声明的所有方法」
//注意： (usb Usb) 体现了多态，根据传入参数类型不一样，它能体现出不同的状态。
func (c Computer) working(usb Usb) { //usb变量会根据传入的实参，判断到底是Phone还是Camera。体现了多态!!!
	//通过usb接口变量来调用
	usb.Start()
	usb.Stop()
}

func interface21() {
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}
	//关键点
	computer.working(phone)
	computer.working(camera)
}
