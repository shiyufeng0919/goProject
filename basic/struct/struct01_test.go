package _struct

import (
	"fmt"
	"testing"
)

/*
「struct是值类型」

一,go语言核心编程-面向对象编程(https://www.bilibili.com/video/BV1ME411Y71o?p=182)

	golang中没有类的概念，使用的是struct.

	一个程序就是一个世界，有很多对象(变量). //创建一个结构体对象(对象也就是变量)

	golang语言面向oop编程的特性:
	1,go也支持面向对象编程（oop），但是和传统面向对象编程有区别，并不是纯粹的面向对象语言。所以说go支持面向对象编程特性是比较准确的.
	2,go没有class,go语言的struct和其它编程语言class有同等地位，可理解go是基于struct来实现oop特性.
	3,go面向对象编程非常简结，去掉了传统oop语言的继承，方法重载，构造函数，析构函数，隐藏的this指针等.
	4,go仍然有面向对象编程的「继承，封装，多态」的特性，只是实现的方式和其它oop语言不一样，如：继承，go没有extends关键字，继承是通过匿名字段来实现.
	5,go面向对象oop很优雅，oop本身氷是语言类型系统(type system)的一部分，通过interface关联，耦合性低，非常灵活。

	「golang中面向接口编程是非常重要的特性!!!」

二。结构体与结构体变量(也叫实例或对象)的关系

	1,将一类事务的特征提取出来,eg: 人类，形成一个新的数据类型，就是一个struct
	2,通过这个struct,可以创建多个变量(实例/对象)
	3,事务可以为person/cat/dog...
	type Person struct{...}
	type Cat struct{...}
	...

三。结构体内存布局(重要!!!) : struct11()

四。声明结构体

	type 标识符 struct { //标识符就是struct的name, struct name首字母大写，则可在其它包被使用
		filed1 type  //字段/属性。 首字母大写,则该字段可被其它包引用
		filed2 type  //字段类型: 基本数据类型/数组/引用类型,若创建一个结构体变量后，未给字段赋值，则默认零值. bool-false,数值-0,字符串->"";指针/slice/map零值为nil.即还未分配空间
	}

五。创建结构体变量和访问结构体字段 -> struct14()

	1,var person Person //直接声明

	2,var person Person = Person{}

	3,var person *Person = new(Person) //返回struct指针

	  结构体指针访问方式(标准): (*结构体指针).字段名  eg:(*Person).Name="syf"

	  go做了简化，也支持 结构体指针.字段名。 如 person.Name = "syf. go编译器底层对person.Name做了转化(*person).Name

	  「注意：(*person).Name不能这样写 *person.Age，理由： .的优先级比*高，会先取person.Age,再取 *」

	4,var person *Person = &Person{} //返回struct指针

六。struct内存分配机制 struct -> 值拷贝

	变量总是存在内存中的，那么结构体变量在内存中是怎样存在的？

	p ->   zs  Name
		   10  Age

	p1 := p  #进行值拷贝,p1会原样复制一份p

	p1 ->  zs  Name
		   10  Age   #修改p1值不会影响p

---------------

	p ->   zs  Name
		   10  Age

	p1 := &p  #将p的地址赋给p1,则p1拥有的是p的地址，p1.Age="12"的改变不仅会改变p1，同时会影响p

    p1 ->  p的地址
*/

func TestStruct01(t *testing.T) {
	//案例, struct是值类型 & 内存布局
	struct11()
	//struct字段类型为指针/slice/map零值为nil.需要先make才能使用
	struct12()
	//不同结构体变量的字段是独立的，互不影响，一个struct变量字段的更改，不影响另外一个。「结构体是值类型」
	struct13()
	//创建结构体变量和访问结构体字段 (4种方式)
	struct14()
}

/*
定义一个Person结构体,给该结构体定义变量(对象/实例)
将Person的特征提取出来，就是变量!!!
为Person增加一些行为动作，如走路，就是方法!!!

结构体是自定义的数据类型，代表一类事务。
struct变量(实例)是具体的，实际的，代表一个具体变量。
*/
type Person struct { //代表一类事务Person
	Name string //字段名大写，可以公开
	Age  int
	Sex  string
}

type Person2 struct {
	Name   string
	Scores [3]float64
	ptr    *int              //指针
	slice  []int             //切片
	map1   map[string]string //map
}

func struct11() {
	/*
		#结构体是值类型，在内在中布局:
		var person Person  //在内存定义一个变量 Person

			person  ---->	 0xc00009a360
								 |
			                 ____""___ Name   ##person声明后的地址
		 					 ____0 ___ Age
							 ____""___ Sex

		##执行person.Name = "syf"...后,修改值。

			person  ---->	 0xc00009a360
								 |
							 ____"syf"___ Name   ##person设置值
							 ____18 _____ Age
							 ____"girl"__ Sex

	*/
	var person Person
	//说明：struct是值类型
	fmt.Printf("struct11===person:%+v,地址:%p\n", person, &person) //struct11===person:{Name: Age:0 Sex:},地址:0xc00009a360
	person.Name = "syf"
	person.Age = 18
	person.Sex = "girl"
	fmt.Printf("struct11===add properties person:%+v ,地址:%p \n", person, &person) //struct11===add properties person:{Name:syf Age:18 Sex:girl} ,地址:0xc000100300
}

func struct12() {
	//struct -> struct是值类型
	var p Person2
	fmt.Println("struct12=== p:", p) //struct12=== p: { [0 0 0] <nil> [] map[]}

	//基本类型
	p.Name = "syf"

	//数组 ->  数组是值类型
	scores := [...]float64{99.0, 88.0, 66.0}
	p.Scores = scores

	//指针类型
	if p.ptr == nil {
		//定义一个int类型， &int取该int地址，即指针类型
		var int = 88
		p.ptr = &int
	}

	//slice切片类型： 使用slice前一定要先make -> 切片是引用类型
	p.slice = make([]int, 2)
	p.slice[0] = 100
	p.slice[1] = 99

	//map数据类型：使用map前一定要先make   -> map是引用类型
	p.map1 = make(map[string]string)
	p.map1["sex"] = "girl"

	fmt.Println("struct12=== set p:", p)    //struct12=== set p: {syf [99 88 66] 0xc00001a308 [100 99] map[sex:girl]}
	fmt.Println("struct12=== ptr:", *p.ptr) //struct12=== ptr: 88

}

func struct13() {
	var p Person
	p.Name = "syf"
	p.Sex = "girl"
	p.Age = 18
	p2 := p
	fmt.Println("struct13===p:", p)   //struct13===p: {syf 18 girl}
	fmt.Println("struct13===p2:", p2) //struct13===p2: {syf 18 girl}
	/*
		p2是值拷备，修改p2不影响p.

		p --->  ___syf___ Name
				___girl__ Sex
				___18____ Age

		p2 ---> ____syf__ Name   #整个值拷备，再修改，不会影响p
				___girl__ Sex
				___18____ Age
	*/
	p2.Age = 35
	fmt.Println("struct13===modify after p:", p)   //struct13===modify after p: {syf 18 girl}
	fmt.Println("struct13===modify after p2:", p2) //struct13===modify after p2: {syf 35 girl}
}

func struct14() {
	var p Person
	//1,
	p.Name = "syf"
	p.Age = 18
	p.Sex = "girl"
	fmt.Println("struct14=== p", p)

	//2,
	p2 := Person{
		Name: "syf",
		Age:  35,
		Sex:  "girl",
	}
	fmt.Println("struct14=== p2", p2)

	//3,定义指针
	var p3 *Person = new(Person)
	//p3是指针，*p3为取值
	(*p3).Name = "syf" // <=> p3.Name
	(*p3).Sex = "girl"
	(*p3).Age = 20
	fmt.Println("struct14=== p3", p3) //struct14=== p3 &{syf 20 girl}
	p3.Age = 11
	fmt.Println("struct14=== modify age p3", p3) //struct14=== modify age p3 &{syf 11 girl}

	//4,
	var p4 *Person = &Person{
		Name: "syf",
		Age:  0,
		Sex:  "girl",
	}
	//注意：此处不能这样写 *p4.Age，理由： .的优先级高，会先取p4.Age,再*
	(*p4).Age = 11
	fmt.Println("struct14=== modify age p4=", p4) //struct14=== modify age p4= &{syf 11 girl}
	p4.Age = 22
	fmt.Println("struct14=== modify age p4:", p4) //struct14=== modify age p4: &{syf 22 girl}
	fmt.Println("struct14=== *p4:", *p4)          //struct14=== *p4: {syf 22 girl}
}
