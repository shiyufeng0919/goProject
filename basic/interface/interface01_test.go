package _interface

import (
	"fmt"
	"syfProject/basic/interface/model"
	"testing"
)

/*
面向对象编程思想-抽象

1,go有面向对象编程的「继承，封装，多态」三大特性。只是实现的方式和其它oop(面向对象)语言不一样。

2,何为抽象
	在定义结构体时，实际上是把一类事物共有属性(字段)和行为(方法)提取出来，形成一个物理模型(模板，即struct)。这种研究问题的方法称为抽象(或编程的一种思维方式).
	张三 -->   不管什么样的人，都有名字，性别，年龄[属性]...   -> type Person struct{...} //物理模型/模板
	李四 -->   都可以走路，吃饭[行为]...							（抽象）
	... -->

3,面向对象编程-封装
	封装(encapsulation)就是把抽象出的「字段和对字段的操作封装」在一起，数据被保护在内部，程序的其它包只有通过被授权的操作(方法)，才能对字段进行操作.
	封装好处:
	1),隐藏实现细节
	2),可以对数据进行验证，保证安全合理(如:age)
	type Person struct{
		Age int //age应在实际范围，在合理情况下，才能设置age
	}
	如何体现封装:
	1),对结构体中的属性进行封装 (如struct字段首字母小写)
	2),通过方法或包来实现封装
	封装的实现步骤:
	step1: 将结构体，字段（属性）的首字母小写(不能导出了，其它包不能使用，类似于private)
	step2: 将结构体所在包提供给一个工厂模式函数，首字母大写。类似于一个构造函数
	step3: 提供一个首字母大写的Set方法(类似于其它语言的public)，用于对属性判断并赋值
		func (var 结构体类型名) SetXxx(参数列表) (返回值列表){
			//加入数据验证的业务逻辑(eg:年龄是否是有效的)
			var.字段=参数
		}
	step4:提供一个首字母大写的Get方法(类似于其它语言的public)，用于获取属性的值
		func (var 结构体类型名) GetXxx(){
			return var.字段
		}
	特别说明：在go开发中并没有特别强调封装，这点不像Java。go本身将面向对象的特性做了简化.

4,面向对象编程-继承
	为什么需要继承？解决代码冗余问题
	如：大学生考试和小学生考试，定义的结构体字段和方法基本一样，但确写了两个几乎相同的代码，这样就出现了代码冗余，并不利于维护和功能扩展。
	即继承可以解决代码复用问题。
	当多个结构体存在相同的属性(字段)和方法时可以从这些结构中抽象出结构体(如Student)，在该结构体中定义这些相同的属性和方法
	其它的结构体不需要重新定义这些属性(字段)和方法，只需要嵌套一个Student匿名结构体即可。
	即:在go中，若「一个struct嵌套了另一个匿名结构体，那么这个结构体可以直接访问匿名结构体的字段和方法，从而实现了继承特性」。
	1).嵌套匿名结构体基本语法
	type Goods struct{
		Name  string
		Price float64
	}
	type Book struct{
		Goods  //这里就是嵌套匿名结构体Goods,继承了Goods的属性(字段)
		Writer string
	}
    2)继承给编程带来的便利
	代码复用性、扩展性、维护性提高了

继承深入讨论:
	1),结构体可以使用嵌套匿名结构体所有的字段和方法，即首字母大写或小写的字段，方法均可以使用。
	2),匿名结构体字段访问
		var book Book
		book.Goods.Name <==> book.Name //编译器先看book对应的类型Book是否有字段Name，若没有，就去看book中嵌入的匿名结构体Goods中有没有Name.若有就调用.
	3),当结构体和匿名结构体有相同的字段或方法时，编译器采用就近访问原则访问。若希望访问匿名结构体的字段和方法，可通过匿名结构体名来区分。
	4),结构体嵌入2/多个匿名结构体，若两个匿名结构体有相同的字段和方法(同时结构体本身没有同名的字段和方法)，在访问时，就必须明确指定匿名结构体名字，否则编译器报错.
	5),若一个struct嵌套了一个有名的结构体，这种模式就是组合，若是组合关系，那么在访问组合的结构体的字段或方式时，必须带上结构体的名字。
		type Goods struct{
			Name  string
			Price float64
		}
		type Book struct{
			goods  Goods //有名的结构体，这种模式就是组合。 Book和Goods组合起来.
			Writer string
		}
		var book Book
		book.Name = "" //错误写法，因为Goods不是匿名结构体，它就不会向上找。因此必须带名字访问。
		book.goods.Name = "" //正确写法
	6),嵌套匿名结构体后，也可以在创建结构体变量(实例)时，直接指定各个匿名结构体字段的值。
		type Goods struct{
			Name  string
			Price float64
		}

		type Brand struct{
			Name    string
			Address string
		}

		//多重继承:嵌套了多个匿名结构体，从而实现了多重继承
		type TV struct{
			Goods  //嵌套匿名结构体
			Brand  //嵌套匿名结构体
		}

		type TV2 struct{
			*Goods  //嵌套匿名结构体
			*Brand
		}
		//可以在创建结构体变量(实例)时，直接指定各个匿名结构体字段的值
		tv1 := TV{Goods{"电视机",5000},Brand{"海尔","山东"}}
		或
		tv2 := TV{Goods{Name: "电视机", Price: 5000,},Brand{Name: "海尔",Address:"山东",}}
		-----
		tv3 := TV{&Goods{"电视机",5000},&Brand{"海尔","山东"}}
		或
		tv4 := TV2{&Goods{Name: "电视机", Price: 5000,},&Brand{Name: "海尔",Address:"山东",}}  //取值: *tv4.Goods 、 *tv4.Brand

	7),结构体的匿名字段是基本数据类型 -> interface13()

5,面向对象编程-多重继承
	如果一个struct嵌套了多个匿名结构体，那么该结构体可以直接访问嵌套的匿名结构体的字段和方法，从而实现了多重继承。
	若嵌入的匿名结构体有相同的字段名或方法名，则在访问时，需要通过匿名结构体类型来区分。
	为了保证代码简结性，建议尽量不使用多重继承.
*/

func TestInterface01(t *testing.T) {
	//封装示例
	interface11()
	//继承示例
	interface12()
	//结构体的匿名字段是基本数据类型，该如何访问？
	interface13()
}

func interface11() {
	//封装
	p := model.NewPerson("syf")
	p.SetAge(18)
	p.SetSalary(5000)
	fmt.Println("interface11===\np:", *p) //p: {syf 18 5000}
}

func interface12() {
	//当对结构体嵌入了匿名结构体后的使用。
	var pupil model.Pupil //或者 pupil := &model.Pupil{}
	fmt.Println("interface12===")
	pupil.Name = "syf" //给model.Pupil{}结构体的Name赋值，而不是给student.Name赋值,所以若调用student.Name，则为""
	pupil.SetAge(18)
	pupil.Test()
	pupil.SetCores(99)
	pupil.ShowInfo()
	pupil.ReturnScore() //结构体可以使用嵌套匿名结构体所有的字段和方法，即首字母大写或小写的字段，方法均可以使用。
	//pupil.Content = "this is a test" //错误,结构体嵌入2/多个匿名结构体，若两个匿名结构体有相同的字段和方法(同时结构体本身没有同名的字段和方法)，在访问时，就必须明确指定匿名结构体名字，否则编译器报错
	pupil.Class.Content = "测试在1个结构体中嵌套多个匿名结构体，且多个匿名结构体中具有相同的字段content.访问问题"

	var graduate model.Graduate //或者 graduate := &model.Graduate{}
	graduate.Name = "smile"     //给student.Name赋值，因为model.Graduate结构体没有Name字段，所以会向上找到匿名结构体student的Name.
	graduate.SetAge(19)
	graduate.Test()
	graduate.SetCores(100)
	graduate.ShowInfo()
}

type Monster struct {
	Name string
	Age  int
}

type Stu struct {
	Monster
	int //基本数据类型也可以用为匿名字段,注意此处有int匿名字段,则不能再有同名int。若要有同名int，需要用名称区分，如n int,m int...
	string
}

func interface13() {
	//演示匿名字段是基本数据类型的使用
	var stu Stu
	stu.Monster.Name = "smile"
	stu.Monster.Age = 500
	stu.int = 100
	stu.string = "syf"
	fmt.Println("interface13===\nstu:", stu)
}
