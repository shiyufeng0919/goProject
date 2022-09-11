package slice

import (
	"fmt"
	"testing"
)

/*
 需求: 需要一个数组用于存储学生成绩，但学生的个数是不确定的。解决方案：使用切片。
 切片: 可理解为动态数组.
	 1.切片是数组的一个引用。因此切片是引用类型。在进行传递时，遵守引用传递的机制。
	 2.切片的使用和数组类似，遍历切片，访问切片的元素和求切片长度len(slice)都一样。
	 3.切片的长度是可以变化的，因此切片是一个可以动态变化的数组。
	 4.切片定义的基本语法
	   var 切片名 []类型   //注意[]不需要添长度
 切片在内存中布局以及切片与其引用的array之间的关系.
*/

func TestSlice01(t *testing.T){
	/*
	1,数组，没有容量概念，是多大就是多大，不能修改.
	2,执行var array [5]int = [5]int{1,2,3,4,5} 会在内存开辟一个数组空间，变量名array指向该空间.
	*/
	var array [5]int = [5]int{1,2,3,4,5}
	fmt.Println("array:",array) //[1 2 3 4 5]
	fmt.Printf("array第0个元素地址:%p \n",&array[0]) //0xc00001c180
	fmt.Printf("array第1个元素地址:%p \n",&array[1]) //0xc00001c188
	/*
	1,slice:切片名称
	2,array[1:3]表示slice引用到array数组下标为[1-3)元素。
	3,执行slice := array[1:3]，会在内存开辟一个空间，第1个位置存储的是array[1]的地址;第2个位置存储的是slice本身的长度；第3个位置存储的是slice的cap。
	  slice指向 -> 「 | array[1]地址 | len(slice) | cap(slice) 」。slice本身也有一个地址。
	  因此slice底层的数据结构实际上是一个struct.
	  struct{
	    ptr *[2]int
	    len int
	    cap int
	  }
	*/
	slice := array[1:3]
	fmt.Println("slice元素:",slice) //[2,3]
	fmt.Println("slice长度:", len(slice)) //2
	/*
	1,cap: 目前可以存放最多个数的元素,切片容量可动态变化。
	2,cap所在包: builtin 内置包中[http://doc.golang.ltd]
	*/
	fmt.Println("slice容量:", cap(slice)) //4 (一般为长度2倍，也不一定.)
	fmt.Printf("slice第1个元素地址:%p \n",&slice[0]) //0xc00001c188

	/*
	1,通过slice修改元素，则array也同时被修改。(因为slice是引用类型，引用的是array数据,slice的改变会影响array)
	*/
	slice[0] = 88
	fmt.Printf("modify slice[0],array:%v,slice:%v",array,slice) //modify slice[0],array:[1 88 3 4 5],slice:[88 3]
}