package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

/*
反射: 最大价值就是写框架

一,反射应用场景
1,应用场景1: 对结构体序列化时，若结构体有指定Tag，也会使用反射生成对应的字符串。
//struct结构体中定义的tag标签使用的是反射机制。
type Person struct{
	Name string `json:"name"
}

2,应用场景2: 不知道接口调用哪个函数，根据传入参数在运行时确定调用的具体接口，这种需要对函数或方法反射。
//定义了两个匿名函数，再定义一个适配器函数，用作统一处理接口。
//anonymous：匿名
anonymousFun1 := func(v1 int, v2 int) {
	fmt.Printf("v1:%v,v2:%v,sum:%v\n", v1, v2, v1+v2)
}
anonymousFun2 := func(v1, v2 int, v3 string) {
	fmt.Printf("v1:%v,v2:%v,v3:%s\n", v1, v2, v3)
}
//定义一个适配器函数，用作统一处理接口。第一个参数funcPtr以接口的形式传入函数指针，函数参数args以可变参数的形式传入.
bridge := func(funcPtr interface{}, args ...interface{}) {
	//内容
}
//实现调用anonymousFun1对应的函数. bridge函数中可以用反射来动态执行funcPtr函数
bridge(anonymousFun1, 88, 99)
//实现调用anonymousFun2对应的函数
bridge(anonymousFun2, 88, 99, "syf")

二,反射基本介绍
1,反射可以在运行时动态获取变量的各种信息，如：变量的类型(type)，类别(kind) {若变量是基本数据类型type和kind是一致的，但若是结构体变量，type和kind不一致。}.
2,如果是结构体变量，还可以获取到结构体本身的信息(包括结构体的字段、方法)
3,通过反射，可以改变变量的值，可以调用关联的方法
4,使用反射，需要import ("reflect")
reflect包实现了运行时反射，允许程序操作任意类型的对象。
典型用法是用静态类型interface{}保存一个值，通过调用TypeOf获取其动态类型信息，该函数返回一个Type类型值。
调用ValueOf函数返回一个Value类型值，该值代表运行时的数据。Zero接受一个Type类型参数并返回一个代表该类型零值的Value类型值。
示意图:参考documents->尚硅谷_韩顺平_Go语言核心编程.pdf
     变量
--------------------
						-----》	reflect.TypeOf  ----》	reflect.Type类型   (操作手册->type Type)
var num int													  |
var person Person								《------------｜ (反向操作)
	...			        -----》	reflect.ValueOf  ----》	reflect.Type类型    (操作手册->type Value)
												《------------|  (反向操作)
---------------------
type Person struct{
	Name string
	Age  int
    ....
}
func (this Person) print(){
//...
}
---------------------
小结:反射重要的函数和概念
1,reflect.TypeOf(变量名),获取变量的类型，返回reflect.Type类型
2,reflect.ValueOf(变量名),获取变量的值，返回reflect.Value类型。reflect.Value是一个结构体类型。
通过reflect.Value可以获取到关于该变量的很多信息。
3,变量、interface{}和reflect.Value是可以相互转换的，在实际开发中，经常使用。
																		     v:类型，为reflect.ValueOf
变量---传递参数---> interface{} ---reflect.ValueOf()函数---> reflect.Value --v.Interface()-->interface{} ----类型断言---> 变量
*/

func TestReflect01(t *testing.T) {
	//编写一个案例，演示对(基本数据类型、interface{}、reflect.Value)进行反射的基本操作。
	reflect11()
	//编写一个案例，演示对(结构体类型、interface{}、reflect.Value)进行反射的基本操作。
	reflect12()
}

func reflect11() {
	var num int = 99
	fmt.Println("reflectTest11===")
	//通过反射获取到传入变量的type、kind、值
	reflectTest11(num)
}

func reflectTest11(i interface{}) {
	//获取到reflect.Type.(注：refType本质上是一个type Type interface{})
	refType := reflect.TypeOf(i)
	fmt.Printf("refType=%v,type=%T,kind=%v\n", refType, refType, refType.Kind()) //refType=int,type=*reflect.rtype,kind=int

	//获取到reflect.Value.(注：refValue本质上是一个type Value struct{})
	refValue := reflect.ValueOf(i)
	fmt.Printf("refValue=%v,type=%T,kind=%v\n", refValue, refValue, refValue.Kind()) //refValue=99,type=reflect.Value,kind=int

	//返回refValue持有的有符号整数（表示为int64），如果refValue的Kind不是Int、Int8、Int16、Int32、Int64会panic
	modifyRefValue := refValue.Int() + 100
	fmt.Printf("modifyRefValue=%v,type=%T\n", modifyRefValue, modifyRefValue) //modifyRefValue=199,type=int64

	//再将refValue转成interface{}
	iValue := refValue.Interface()
	//将interface再通过类型断言转成具体需要的类型
	num := iValue.(int)
	fmt.Printf("num=%v\n", num) //num=99
	fmt.Println("----------reflectTest11 End.-------------")
}

type Student struct {
	Name string
	Age  int
}

type Person struct {
	Name string
	Age  int
}

func reflect12() {
	var stu Student
	stu.Name = "syf"
	stu.Age = 18
	reflectTest12(stu)

	fmt.Println("-----------------------------------------")

	var person Person
	person.Name = "smile"
	person.Age = 19
	reflectTest12(person)
	fmt.Println("----------reflectTest12 End.-------------")
}

func reflectTest12(i interface{}) {
	//获取到reflect.Type
	refType := reflect.TypeOf(i)
	fmt.Printf("refType=%v,kind=%v\n", refType, refType.Kind()) //refType=reflect.Student,kind=struct

	//获取到reflect.value
	refValue := reflect.ValueOf(i)
	fmt.Printf("refValue=%v,type=%T,kind=%v\n", refValue, refValue, refValue.Kind()) //refValue={syf 18},type=reflect.Value,kind=struct

	//将reflect.value转成interface{}
	iValue := refValue.Interface()
	fmt.Printf("iValue=%v,Type=%T\n", iValue, iValue) //iValue={syf 18},Type=reflect.Student
	//使用类型断言转成具体需要的类型.注意此处，如何知道要转成Student还是Person结构体类型，可用下述方法做验证.
	switch iValue.(type) {
	case Student:
		stu, ok := iValue.(Student)
		if ok {
			fmt.Printf("stu:%+v,stu.Name:%s\n", stu, stu.Name) //stu:{Name:syf Age:18},stu.Name:syf
		}
	case Person:
		person, ok := iValue.(Person)
		if ok {
			fmt.Printf("person:%+v,person.Name:%s\n", person, person.Name)
		}
	default:
		fmt.Println("unKnow type.")
	}
}
