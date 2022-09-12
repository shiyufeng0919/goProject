package _struct

import (
	"fmt"
	"syfProject/basic/struct/model"
	"testing"
)

/*
方法:
1,go中的方法是作用在指定的数据类型上的，即和指定的数据类型绑定。因此自定义类型，都可以有方法，而不仅仅是struct.
func (p People) walk(){...} //方法,p People体现walk()方法是和People绑定的。只能通过People类型的变量来调用，而不能直接调用，也不能使用其它类型变量调用

2,方法的调用和传参机制原理(重要!!!)
方法的调用和传参机制和函数基本一样，差异点是方法调用时，会将调用方法的变量，当做实参也传递给方法。
如果变量是值类型，则进行值拷贝，若变量是引用类型，则进行地址拷贝.

3,方法的声明(定义)
//receiver type表示这个methodName和type这个类型进行绑定，或者说该方法作用于type类型。
//type可以是struct，也可以是其它自定义类型；receiver就是type类型的一个变量(实例)
//eg: p Person，则type是Person，p是type类型的一个变量/实例
func (receiver type) methodName(params list)(返回值列表){
	method body
    return 返回值
}

4,方法的注意事项和细节
1),struct类型是值类型，在方法调用时，遵守值类型的传递机制，是值拷贝传递方式.
2),若希望在方法中，修改结构体变量的值，可通过结构体指针方式.
3),go中的方法作用在指定的数据类型上(即和指定的数据类型绑定)，因此自定义类型，都可以有方法，而不仅是struct.如int,float32等都可以有方法.
4),方法的访问范围控制的规则，和函数一样。方法名首字母小写，只能在本包访问，首字母大写，可跨包访问.
5),若一个类型实现了String()这个方法，则fmt.Println默认会调用这个类型的String()进行输出.

5,方法和函数区别

1),调用方式不一样
函数: 函数名(实参列表)
方法: 变量.方法名(实参列表)
2),普通函数，接收者为值类型时，不能将指针类型的数据直接传递。反之亦然。
3),方法(如:struct方法)，接收者为值类型时，可以直接用指针类型的变量调用方法。反之亦然。
总结: 不管调用形式如何，真正决定是值拷贝还是地址拷贝，看这个方法和哪个类型绑定。若是和值类型绑定，如(p Person)则是值拷贝，若是和指针类型，如(p *Person)，则是地址拷贝。

6,面向对象编程应用实例
步骤:
step1: 声明/定义结构体，确认结构体名
step2: 编写结构体字段
step3: 编写结构体方法

7,工厂模式
golang的结构体没有构造函数，通常可以使用工厂模式来解决这个问题.即工厂模式相当于构造函数的功能。
如:一个struct声明
package model
type Student struct{
	Name string...
}
因为Student的首字母大写，因此在其它包(如:main)引入model包即可直接创建Student结构体的变量/实例。
但问题是，若首字母小写(需求)，则该如何解决？--工厂模式
即：使用工厂模式解决跨包访问结构体首字母小写或字段小写问题!!!
*/
func TestStruct03(t *testing.T) {
	//go中的方法是作用在指定的数据类型上的，即和指定的数据类型绑定。因此自定义类型，都可以有方法，而不仅仅是struct.
	struct31()
	//方法的调用和传参机制原理
	struct32()
	//方法的调用和传参机制原理
	//struct类型是值类型，在方法调用时，遵守值类型的传递机制，是值拷贝传递方式.
	//若希望在方法中，修改结构体变量的值，可通过结构体指针方式.
	struct33()
	//go中的方法作用在指定的数据类型上(即和指定的数据类型绑定)，因此自定义类型，都可以有方法，而不仅是struct.如int,float32等都可以有方法.
	struct34()
	//若一个类型实现了String()这个方法，则fmt.Println默认会调用这个类型的String()进行输出.
	struct35()
	//普通函数，接收者为值类型时，不能将指针类型的数据直接传递。
	//方法(如:struct方法)，接收者为值类型时，可以直接用指针类型的变量调用方法。
	struct36()
	//工厂模式，解决跨包创建结构体实例案例.
	struct37()
}

type People struct {
	Name string
	Age  int
}

type Circle struct {
	Radius float64
}

type Students struct {
	Name string
	Age  int
}

//这是一个函数
func struct31() {
	var p People
	p.Name = "syf"
	p.Age = 18
	p.walk()                                 //调用方法,必须是People类型的变量来调用
	fmt.Println("struct31===\nwalk() p:", p) //p: {syf 18}
	p.run()
	fmt.Println("struct31===\nrun() p:", p) //p: {syf 36}
}

func struct32() {
	var p People
	p.Name = "syf"
	p.Age = 18
	n1 := 10
	n2 := 20
	/*
		内存图:
		--------------
		struct32栈

		n1  -> [10]
		n2  -> [20]
		p   -> struct[syf] Name ...   //结构体值拷贝
		sum := 30

		--------------

		getSum栈

		n1  ->  [10]
		n2  ->  [20]
		p   -> struct[syf ] Name ...   //结构体值拷贝，完整拷贝一份

		//计算n1+n2后return struct32栈,sum接收...
	*/
	sum := p.getSum(n1, n2)
	fmt.Println("struct32===\nsum:", sum) //sum: 30
}

func struct33() {
	/*
		内存结构图
		------------------------------------------------
		  struct33栈

		  c -> [4.0] radius

		  res:= 3.14
		------------------------------------------------
		  area栈

		  c -> [1.0] radius

		  运算3.14 * radius * radius，并返回struct33栈res，此时area栈该值被销毁.
		------------------------------------------------
	*/
	//方式一：值拷贝
	var c Circle
	c.Radius = 4.0
	res := c.area()
	fmt.Printf("struct33===\nres:%v,radius:%v \n", res, c.Radius) //res:3.14,radius:4

	/*
			内存结构图
			------------------------------------------------
			  struct33栈

		        0xc0001261f8
			  c -> [4.0] radius

			  res:= 3.14
			------------------------------------------------
			  area栈

			  c -> [0xc0001261f8] radius

		      执行 (*c).Radius = 1.0 则实质为修改了0xc0001261f8指向地址(struct33栈)的radius值

			  运算3.14 * radius * radius，并返回struct33栈res，此时area栈该值被销毁.
			------------------------------------------------
	*/
	//方式二：引用拷贝
	var c2 Circle
	fmt.Printf("c2 Circle address:%p \n", &c2) //c2 Circle address:0xc0001261f8
	c2.Radius = 4.0
	res2 := (&c2).areas()                               //标准访问 <=> res2 := c2.areas()
	fmt.Printf("res2:%v,radius:%v \n", res2, c2.Radius) //res2:3.14,radius:1
}

func struct34() {
	var i integer
	fmt.Println("===struct34")
	i = 10
	i.print()
	fmt.Println("===struct34 i:", i) //===struct34 i: 10

	var i2 integer
	i2 = 10
	i2.change()
	fmt.Println("===struct34 i2:", i2) //===struct34 i2: 20
}

func struct35() {
	//定义一个Students变量
	var stu Students
	stu.Name = "syf"
	stu.Age = 18
	//实现了*Students类型的String方法，就会自动调用.若未实现，则打印的是Students地址
	fmt.Println("===struct35 stu:", &stu) //===struct35 stu: name=[syf],age=[18]
}

func struct36() {
	p := People{Name: "syf", Age: 18}
	fmt.Println("===struct36")

	//printPerson()函数接收的是值，则此处只能传值，不能传地址.[必须类型匹配]
	printPerson(p)
	//printPerson()函数接收的是地址，则此处只能传地址，不能传值.[必须类型匹配]
	printPerson1(&p)

	//下述两种方式均可以，均是「值拷贝」,因为printPerson2()定义的type是值类型!!!
	p.printPerson2()
	fmt.Println("===struct36 p.name1:", p.Name) //===struct36 p.name1: syf
	(&p).printPerson2()
	fmt.Println("===struct36 p.name2:", p.Name) //===struct36 p.name2: syf

	//下述两种方式均可以，均是「引用拷贝」,因为printPerson2()定义的type是指针类型!!!
	p.printPerson3()
	fmt.Println("===struct36 p.name3:", p.Name) //===struct36 p.name3: smile
	(&p).printPerson3()
	fmt.Println("===struct36 p.name4:", p.Name) //===struct36 p.name4: smile
}

func struct37() {
	//跨包创建Student实例
	var stu = model.Student{
		Name: "syf",
		Age:  18,
	}
	fmt.Println("===struct37\nstu:", stu) //stu: {syf 18}

	//使用工厂模式解决跨包访问首字母小写struct实例的方案
	score := model.NewScores("syf", 100)
	level := score.GetScoreLevel()
	fmt.Printf("===struct37\nscore:%+v,level:%s \n", *score, level) //score:{Name:syf Score:100 level:middle},level:middle
}

/*
这是一个方法
1,方法walk()和People类型绑定.
2,walk()方法只能通过People类型的变量来调用，而不能直接调用，也不能使用其它类型变量调用。
*/
func (p People) walk() {
	p.Age = 35                                            //内部的改变不会影响walk()外部
	fmt.Printf("walk(),name:%s,age:%v \n", p.Name, p.Age) //walk(),name:syf,age:35
}

/*
这是一个方法
*/
func (p *People) run() {
	p.Age = 36                                           //内部改变会影响run()外部
	fmt.Printf("run(),name:%s,age:%v \n", p.Name, p.Age) //run(),name:syf,age:36
}

func (p *People) getSum(n1, n2 int) int {
	return n1 + n2
}

//返回面积
func (c Circle) area() float64 {
	c.Radius = 1.0 //修改值
	return 3.14 * c.Radius * c.Radius
}

func (c *Circle) areas() float64 {
	fmt.Printf("c是*Circle指向的地址:%p \n", c) //c是*Circle指向的地址:0xc0001261f8
	(*c).Radius = 1.0
	//<=> 3.14 * c.Radius * c.Radius
	return 3.14 * (*c).Radius * (*c).Radius //标准用法
}

//值传递
func (i integer) print() {
	fmt.Println("i=", i) //10
}

//引用传递
func (i *integer) change() {
	*i = *i + 10          //方法内修改值，会影响方法外
	fmt.Println("i=", *i) //20
}

//给Students实现String()方法
func (stu *Students) String() string {
	str := fmt.Sprintf("name=[%s],age=[%v]", stu.Name, stu.Age)
	return str
}

func printPerson(p People) {
	fmt.Println(p.Name)
}

func printPerson1(p *People) {
	fmt.Println((*p).Name)
}

//值拷贝
func (p People) printPerson2() {
	p.Name = "smile" //修改此值，不影响方法外
	fmt.Println(p.Name)
}

//引用拷贝
func (p *People) printPerson3() {
	p.Name = "smile" //修改此值，则影响方法外
	fmt.Println(p.Name)
}
