package grammar

import (
	"fmt"
	"testing"
)

/*
Const常量

一，常量介绍：
1，常量使用const修饰
2，常量在定义的时候必须初始化
3，常量不能修改
4，常量只能修饰bool,数值类型(int、float系列)、string类型
5，语法：const identifier [type] = value
6，举例：
(1),const name = "syf"
(2),const tax float64 = 0.8
(3),const a int           //错误，常量在定义的时候必须初始化
(4),const b = 9/3
(5),const c = getValue()  //错误，常量不能修改，值必须是一个确定value。

二，常量使用注意事项：
1,简洁写法
const (
	a = 1
    b = 2
)
2,专业写法
const (
	a = iota   //表示给a赋值为0;b在a的基础上+1;c在b的基础上+1
	b
	c
)
3,golang中没有常量名必须字母大写的规范，如:TAX_RATE
4,仍然通过首字母的大小写来控制常量的访问范围
*/

func TestConst(t *testing.T) {
	//注意下述常量用法的特别
	const01()
}

func const01() {
	const (
		a = iota
		b
		c
		d
	)
	fmt.Println(a, b, c, d) //0 1 2 3

	const (
		a1     = iota
		b1     = iota
		c1, d1 = iota, iota
	)
	fmt.Println(a1, b1, c1, d1) //0 1 2 2

	const (
		a2 = iota
		b2 = iota
		c2 = iota
		d2 = iota
	)
	fmt.Println(a2, b2, c2, d2) //0 1 2 3
}
