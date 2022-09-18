package _interface

import (
	"fmt"
	"testing"
)

/*
一。多态及实现
1，基本介绍
变量(实例)具有多种形态。面向对象第三大特征，在go语言中，「多态特征是通过接口实现的」。
可以按照统一的接口来调用不同的实现。这时接口变量就呈现不同的形态。
例： USB接口示例。interface02_test()
Usb usb,即可以接收手机变量，又可以接收相机变量。就体现了Usb接口多态特性。

2,接口体现多态的两种形式
1),多态参数
   Usb usb,即可以接收手机变量，又可以接收相机变量。就体现了Usb接口多态特性。
2),多态数组
 给Usb数组，存放Phone结构体和Camera结构体变量。
 Phone还有一个特有的方法call()，请遍历usb数组，如果是phone变量，除了调用Usb接口声明的方法外，还需要调用Phone特有方法call.(需要使用类型断言)


二。类型断言

1，类型断言，由于接口是一般类型，不知道具体类型，若要转成具体类型，就需要使用类型断言。

2,类型断言最佳实践
1)，在前端Usb接口案例做改进:给Phone结构体增加一个特有方法call()，当Usb接口接收的是Phone变量时，还需要调用Call方法

2), 写一个函数，循环判断传入参数的类型
*/
func TestInterface07(t *testing.T) {
	//给Usb数组，存放Phone结构体和Camera结构体变量。
	interface71()
	//类型断言(如何将一个接口变量，赋给自定义类型的变量 => 引出类型断言)
	interface72()
	//类型断言其它案例：在进行类型断言时，若类型不匹配，则panic，因此进行类型断言时，要确保类型原来的空接口指向的就是要断言的类型。
	interface73()
	//如何在类型断言时，带上检测机制，若成功则ok，否则也不要报panic.
	interface74()
	//在前端Usb接口案例做改进:给Phone结构体增加一个特有方法call()，当Usb接口接收的是Phone变量时，还需要调用Call方法
	interface75()
	//写一个函数，循环判断传入参数的类型(应用可变参数...interface{}来接收任意多个类型)
	//类型断言不仅可以断言已有的数据类型，也可以断言自定义的类型。
	interface76()
}

func interface71() {
	//定义一个Usb接口数组，可以存放Phone和Camera的结构体变量。这里就体现了多态数组
	var usbArray [3]Usb
	fmt.Println(usbArray) //[<nil> <nil> <nil>]

	//利用接口多态特性，一个数组可以放不同类型的值
	usbArray[0] = Phone{}
	usbArray[1] = Phone{}
	usbArray[2] = Camera{}
	fmt.Println(usbArray) //[{} {} {}]
}

type Point struct {
	x int
	y int
}

func interface72() {
	var a interface{} //定义一个空接口
	var point Point = Point{x: 8, y: 9}
	// point -> a
	a = point //空接口可以接收任意类型变量,「a指向Point类型变量」

	//如何将a赋给一个Point变量?
	var b Point
	//b = a  //错误写法: 不能将a作为一个Point类型直接赋给b

	// a -> point 将interface重新转成Point
	b = a.(Point)  //判断「a是否指向Point类型变量」，如果是就转成Point类型并赋给b变量，否则报错。此种写法即为「类型断言」
	fmt.Println(b) //{8 9}
}

func interface73() {
	var x interface{} //定义一个空接口
	var b float32 = 8.9
	//b -> x
	x = b //空接口可以接收任意类型变量,「a指向b类型变量」
	//x -> b
	b = x.(float32)                                //将interface{}重新转成float32
	fmt.Printf("interface73===\n x:%v,b:%v", x, b) // x:8.9,b:8.9
}

func interface74() {
	var a interface{} //定义一个空接口
	var point Point = Point{x: 8, y: 9}
	// point -> a
	a = point //空接口可以接收任意类型变量,「a指向Point类型变量」

	//如何将a赋给一个Point变量?
	var b Point
	//b = a  //错误写法: 不能将a作为一个Point类型直接赋给b

	// a -> point 将interface重新转成Point
	if b2, ok := a.(Point); ok { //判断「a是否指向Point类型变量」，如果是就转成Point类型并赋给b变量，否则报错。此种写法即为「类型断言」
		fmt.Printf("interface74===\nconvert success,b2的类型是：%T，值是:%+v \n", b2, b2) //convert success,b2的类型是：_interface.Point，值是:{x:8 y:9}
		b = b2
		fmt.Println(b) //{8 9}
	} else {
		fmt.Println("convert fail")
		return
	}
}

//属于Phone结构体独有方法
func (p Phone) Call() {
	fmt.Println("phone call work.")
}

type Digital struct {
}

func (d Digital) working(usb Usb) {
	usb.Start()

	//改动：当Usb接口接收的是Phone变量时，还需要调用Call方法 「类型断言」
	if phone, ok := usb.(Phone); ok { //断言成功
		phone.Call()
	}

	usb.Stop()
}

//基于interface02_test.go
func interface75() {
	var usbArray [3]Usb
	usbArray[0] = Phone{"HuaWei"}
	usbArray[1] = Phone{"Apple"}
	usbArray[2] = Camera{"Nero"}
	var digital Digital
	for _, v := range usbArray {
		digital.working(v)
	}
}

func interface76() {
	fmt.Println("interface76===")
	var n1 float32 = 9.8
	var n2 float64 = 8.9
	var n3 int32 = 100
	var n4 string = "syf"
	n5 := "smile"
	n6 := 999
	point := Point{}
	TypeJudge(n1, n2, n3, n4, n5, n6, point, &point)
}

//...表示可变参数，表示可以接收任意多个类型的实参
func TypeJudge(items ...interface{}) {
	for k, v := range items {
		switch v.(type) { //类型断言
		case bool:
			fmt.Printf("param %d is a bool,value is :%v\n", k, v)
		case float64:
			fmt.Printf("param %d is a float64,value is :%v\n", k, v)
		case int, int64:
			fmt.Printf("param %d is an int,value is :%v\n", k, v)
		case nil:
			fmt.Printf("param %d is nil,value is :%v\n", k, v)
		case string:
			fmt.Printf("param %d is a string,value is :%v\n", k, v)
		case Point:
			fmt.Printf("param %d is a Point,value is :%v\n", k, v)
		case *Point:
			fmt.Printf("param %d is a *Point,value is :%v\n", k, v)
		default:
			fmt.Printf("param %d type unknown,value is :%v\n", k, v)
		}
	}
}
