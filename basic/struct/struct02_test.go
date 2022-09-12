package _struct

import (
	"encoding/json"
	"fmt"
	"testing"
)

/*
一。结构体注意事项和使用细节
1,struct的所有字段在「内存」中是连续的
2,struct是用户单独定义的类型，与其它类型进行转换时需要有完全相同的字段(名字，个数，类型必须完全相同)
3,struct进行type重新定义（相当于取别名），go认为是新的数据类型，但是相互间可强转.
4,struct的每个字段上，可以写上一个tag,该tag通过反射机制获取，常见：序列化和反序列化。
*/

type Point struct {
	x int
	y int
}

type Rect struct {
	leftUp, rightDown Point //存储具体Point值
}

type Rect2 struct {
	leftUp, rightDown *Point //存储Point地址
}

func TestStruct02(t *testing.T) {
	//struct的所有字段在「内存」中是连续的
	struct21()
	//struct是用户单独定义的类型，与其它类型进行转换时需要有完全相同的字段(名字，个数，类型必须完全相同)
	struct22()
	//struct进行type重新定义（相当于取别名），go认为是新的数据类型，但是相互间可强转.
	struct23()
	//struct的每个字段上，可以写上一个tag,该tag通过反射机制获取，常见：序列化和反序列化。
	//将struct变量进行json处理
	struct24()
}

func struct21() {
	//r中有4个int,在内存中是连续分布的
	r := Rect{Point{1, 2}, Point{3, 4}}
	/*
		  r ->  Rect struct {}
				0xc000016120  0xc000016128   0xc000016130   0xc000016138
				|   1      |        2       |  3        |         4    |
	*/
	//r.leftUp.x:0xc000016120,r.leftUp.y:0xc000016128,r.rightDown.x:0xc000016130,r.rightDown.y:0xc000016138
	fmt.Printf("struct21===\n r.leftUp.x:%p,r.leftUp.y:%p,r.rightDown.x:%p,r.rightDown.y:%p \n", &r.leftUp.x, &r.leftUp.y, &r.rightDown.x, &r.rightDown.y)

	/*
		  r2  -> Rect2 struct{}
				0xc000118500			   0xc000118508				   ->本身地址
			|  地址1(0xc00001a2e0)    |    地址2(0xc00001a2f0)    |     ->指向地址
	*/
	//r2有两个*Point类型，这2个*Point类型本身地址也是连续的，但他们指向的地址不一定是连续的
	r2 := Rect2{&Point{10, 20}, &Point{30, 40}}
	//r2.leftUp本身地址:0xc000118500,r2.rightDown本身地址:0xc000118508
	fmt.Printf("r2.leftUp本身地址:%p,r2.rightDown本身地址:%p \n", &r2.leftUp, &r2.rightDown)
	//上述指向的地址不一定是连续的,这要看系统在运行时是如何分配的.
	//r2.leftUp指向地址:0xc00001a2e0,r2.rightDown指向地址:0xc00001a2f0
	fmt.Printf("r2.leftUp指向地址:%p,r2.rightDown指向地址:%p \n", r2.leftUp, r2.rightDown)
}

type A struct {
	Name int
}

type B struct {
	Name int
}

func struct22() {
	var a A
	var b B
	//a = b  //错误：不能直接将b类型赋值给a
	a = A(b)                                         //将b类型强制转换为A,可以.前提A和B类型的字段「名字，类型及个数」必须一致。
	fmt.Printf("struct22===\n a:%+v,b:%+v \n", a, b) // a:{Name:0},b:{Name:0}
}

type Student struct {
	Name string
	Age  int
}

type Stu Student //struct进行type重新定义（相当于取别名)

type integer int //struct进行type重新定义，将int取别名为integer

func struct23() {
	var stu1 Student
	var stu2 Stu
	//stu2 = stu1 //错误,两个不相同的类型
	stu2 = Stu(stu1)                                             //将stu1强转为stu2类型Stu
	fmt.Printf("struct23=== \nstu1:%+v,stu2:%+v \n", stu1, stu2) //stu1:{Name: Age:0},stu2:{Name: Age:0}

	var i integer = 10
	var j int = 20
	//j = i //错误，上述使用type将int重新定义为integer,则integer和int被视为两个不相同的数据类型
	j = int(i)                       //将i变量的类型integer强制转换成j类型int
	fmt.Printf("i:%v,j:%v \n", i, j) //i:10,j:10
}

type Monster struct {
	Name string `json:"name"` // json:"xxx" 就是struct的tag.
	Age  int    `json:"age"`
}

func struct24() {
	m := Monster{
		Name: "syf",
		Age:  18,
	}
	fmt.Printf("struct24===\nm:%+v \n", m) //m:{Name:syf Age:18}
	/*  http://doc.golang.ltd -> encoding/json
	将m变量序列化为json字符串
	注意：Monster及其字段均需要首字母大写，理由json在其它包，若想要能够访问到Monster及其字段，则必须大写。
	*/
	jsonM, err := json.Marshal(m) //json.Marshal应用到反射
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("jsonM:%s \n", string(jsonM)) //jsonM:{"name":"syf","age":18}
}
