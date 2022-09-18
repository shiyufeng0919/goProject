package _interface

import (
	"fmt"
	"testing"
)

/*
面向对象编程思想-接口  「interface是引用类型」
一。基本介绍
interface类型可以定义一组方法，但是这些不需要实现。并且interface「不能包含任何变量」。
到某个自定义类型(eg:结构体Phone)要使用的时候，在根据具体情况把这些方法写出来。

二。基本语法
type 接口名 interface{			   「实现接口所有方法」			method1(params list)函数签名必须与接口中定义一致!!!
	method1(params list)返回值列表   ------------>  func (t 自定义类型) method1(params list) return list{ //method implement}
	method2(params list)返回值列表				  func (t 自定义类型) method2(params list) return list{ //method implement}
	...												...
}
小结说明:
1,接口里的所有方法都没有方法体，即接口的方法都是没有实现的方法。接口体现了程序设计的多态和高内聚低耦合的思想。
2,go中的接口，不需要显示的实现。「只要一个变量，含有接口类型中的所有方法，那么这个变量就实现这个接口」(即:go中的接口是基于方法的)。因此
  go中没有implement这样的关键字。

三。应用场景
1，什么时候使用接口？    项目经理「定规范」，其它程序员「执行」

四。注意事项和细节
1,接口本身不能创建实例，但是可以指向一个实现了该接口的自定义类型的变量(实例)
2,接口中所有的方法都没有方法体，即都是没有实现的方法。
3,在go中，一个自定义类型需要将某个接口的所有方法都实现，我们说这个自定义类型实现了该接口。
4,一个自定义类型只有实现了某个接口，才能将该自定义类型的实例(变量)赋值给接口类型。
5,「只要是自定义数据类型，就可以实现接口，不仅仅是结构体类型」.
6,一个自定义类型可以实现多个接口。
7,go接口中不能有任何变量。
8,一个接口(eg:A接口)可以继承多个别的接口(eg:B，C接口)，这时若要实现A接口，则也必须将B，C接口的方法也全部实现.
9,interface类型默认是一个指针，如果没有对interface初始化就使用，那么会输出nil。 「interface是引用类型」
10,空接口interface{}没有任何方法，所以所有类型都实现了空接口。「即我们可以把任何一个变量赋给空接口类型」
*/

func TestInterface03(t *testing.T) {
	//接口本身不能创建实例
	//interface31()
	//接口本身不能创建实例，但是可以指向一个实现了该接口的自定义类型的变量(实例)
	interface32()
	//只要是自定义数据类型，就可以实现接口，不仅仅是结构体类型
	interface33()
	//一个自定义类型可以实现多个接口
	interface34()
	//一个接口(eg:A接口)可以继承多个别的接口(eg:B，C接口)，这时若要实现A接口，则也必须将B，C接口的方法也全部实现.
	interface35()
	//空接口interface{}没有任何方法，所以所有类型都实现了空接口。即我们可以把任何一个变量赋给空接口类型。
	interface36()

	//interface{}重复定义方法名，导致编译通不过
	interface37()
	//接口由指针类型实现
	interface38()
}

type AInterface interface {
	Say()
}

type student struct {
}

//实现了AInterface
func (stu student) Say() {
	fmt.Println("Say syf...")
}

func interface31() {
	fmt.Println("interface31===")
	var a AInterface //接口本身不能创建实例
	a.Say()          //panic: runtime error: invalid memory address or nil pointer dereference [recovered]
}

func interface32() {
	fmt.Println("interface32===")
	var stu student        //结构体变量stu实现了Say()实现了AInterface
	var a AInterface = stu //接口本身不能创建实例，但是可以指向一个实现了该接口的自定义类型的变量stu
	a.Say()                //Say syf...
}

//只要是自定义数据类型，就可以实现接口，不仅仅是结构体类型
type integer int

func (i integer) Say() {
	fmt.Println("Say smile...")
}
func interface33() {
	fmt.Println("interface33===")
	var i integer = 10
	var a AInterface = i
	a.Say() //Say smile...
}

type BInterface interface {
	Hello()
}

//一个自定义类型可以实现多个接口
type stu struct {
}

//stu实现了AInterface
func (s stu) Say() {
	fmt.Println("stu say...")
}

//stu实现了BInterface
func (s stu) Hello() {
	fmt.Println("stu hello...")
}

func interface34() {
	fmt.Println("interface34===")
	var s stu
	var a AInterface = s //将s变量给A接口
	var b BInterface = s //将s变量给B接口
	a.Say()              //stu say...
	b.Hello()            //stu hello...
}

type CInterface interface {
	Ok()
	AInterface //继承了AInterface
	BInterface //继承了BInterface
}

//若要实现CInterface，则需要将AInterface,BInterface及Ok()方法全部实现
type smile struct {
}

func (s smile) Ok() {
	fmt.Println("smile ok...")
}

func (s smile) Say() {
	fmt.Println("smile say...")
}

func (s smile) Hello() {
	fmt.Println("smile hello...")
}

func interface35() {
	fmt.Println("interface35===")
	var s smile
	var a AInterface = s
	var b BInterface = s
	var c CInterface = s
	a.Say()   //smile say...
	b.Hello() //smile hello...
	c.Ok()    //smile ok...

	var s2 smile
	var c2 CInterface = s2
	c2.Ok()    //smile ok...
	c2.Say()   //smile say...
	c2.Hello() //smile hello...
}

type T interface {
}

func interface36() {
	//我们可以把任何一个变量赋给空接口类型。
	fmt.Println("interface36===")
	var s stu
	var t T = s
	fmt.Println(t) //{}

	var t2 interface{} //定义t2是一个空接口类型
	t2 = s
	fmt.Println(t2) //{}
	var num float64 = 9.9
	t2 = num
	fmt.Println(t2) //9.9
}

type A interface {
	test01()
	test02()
}

type B interface {
	test01()
	test03()
}

//错误用法：C继承接口A和B，有两个相同test01()方法，编译不能通过.
//type C interface {
//	A
//	B
//}

func interface37() {
	//接口C继承A，B接口，由于A有test01(),B也有test01()，导致C的interface中报重复定义，不能通过.
}

type D interface {
	Say()
}

type E struct {
}

//D接口，是由指针类型实现的
func (e *E) Say() {
	fmt.Println("E Say...")
}

func interface38() {
	fmt.Println("interface38===")
	var e E
	var d D = &e //D接口，是由指针类型实现的，所以此处要传地址.
	d.Say()
}
