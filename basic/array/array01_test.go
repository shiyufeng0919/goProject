package array

import (
	"fmt"
	"testing"
)

/*
「数组是值类型」

1.数组可以存放多个"同一类型"数据,数组元素index从 0 -> n。
2.数组也是一种数据类型，在GO中，数组是“值类型”。
3.数组在内存中的布局.数组在内存中的地址是连续的。
5.%p: p->pointer 即%p取出地址。
6.数组的4种声明方式、遍历数组。
7.注意事项和使用细节
(1).数组是多个相同类型数据的组合，一个数组一旦声明/定义了，其长度是固定的，不能动态变化。(否则数组越界)
(2).var arr []int 这时arr是一个slice切片，而不是数组
(3).数组中的元素可以是任何数据类型，包括值类型和引用类型，但不能混用。 如var array[2] int只能放int类型不能放其它.
(4).数组创建后，若没有赋值，有默认值(零值)。数值类型数组默认值为0，字符串为"", bool为false
(5).使用数组步骤: 声明数组并开辟空间；为数组各个元素赋值（默认零值）；使用数组
(6).数组的下标是从0开始
(7).数组下标必须在指定范围内使用，否则panic: 数组越界。
(8).Go的数组数据类型，默认情况下是值传递。因此会进行值拷贝。数组间不会相互影响。
(9).如想在其它函数中，去修改原数组，可使用引用传递（指针方式）
(10).长度是数组类型的一部分，在传递函数参数时，需要考虑数组的长度。即不能把var arr = [...]{1,2,3}传给func test(arr [2]int)
另：不能把数组var arr = [...]{1,2,3} 传给func test(arr int) //此处arr为切片；
*/
func TestArray01(t *testing.T) {
	//定义数组1
	var array [2]float64             //在内存开辟3个空间，每个空间值为0。array指向该数组空间的首地址(即array[0]地址)。数组的地址可以通过数组名获取.且数组的首地址即为数组的第一个元素地址.
	fmt.Printf("数组地址:%p \n", &array) //0xc00001a2d0
	array[0] = 1.0
	array[1] = 2.0                                               //float64是8个字节(64位，每8位为1个字节；即64/8=8个字节)；float32是4个字节
	fmt.Printf("第1个元素地址:%p;第2个元素地址:%p \n", &array[0], &array[1]) //0xc00001a2d0,0xc00001a2d8 (第2个元素地址=第1个元素地址+第2个元素所占的字节数)
	//遍历数组1
	for i := 0; i < len(array); i++ {
		fmt.Printf("i:%v , i/array[i]:%s \n\n", i, fmt.Sprintf("%.2f", float64(i)/array[i]))
	}

	//4种初始化数组方式
	var array1 [2]int = [2]int{1, 2} //指定类型
	fmt.Println("array1:", array1)

	var array2 = [2]int{3, 4} //类型推倒
	fmt.Println("array2:", array2)

	var array3 = [...]int{5, 6} //不指定数组大小
	fmt.Println("array3:", array3)

	var array4 = [...]int{1: 888, 0: 999, 2: 666} //指定元素下标给值
	fmt.Println("array4:", array4)                // [999 888 666]

	//遍历数组2
	for k, v := range array4 {
		//k:为数组的下标；v:为数组下标对应的值 （k,v为自定义，仅在for内部有效-局部变量）
		//若只想要v,不想要k,则k可用_。如 for _,v := range array4{}
		fmt.Printf("k:%v,v:%v \n", k, v)
		fmt.Printf("array4[%d] = %v \n", k, array4[k])
	}

	//Go的数组数据类型，默认情况下是值传递。因此会进行值拷贝。数组间不会相互影响。
	modifyArray(array1)                                              //array1在内存中已开辟一个数组空间
	fmt.Printf("array value copy,modify array1[0]:%v \n", array1[0]) //1

	//使用指针传递，达到在另一function中修改数组的目的
	modifyArrayByPoint(&array1)
	fmt.Printf("array point copy,modify array1[0]:%v \n", array1[0]) //689
}

//此处arr为值拷贝，会再开辟一个新的栈空间，在其内部的所有改变，都只是在改变新空间的值，而不会影响原空间的值
//效率低（值拷备）
func modifyArray(arr [2]int) { //Go中[2]int长度数据类型一部分，必须写上，若不写长度，则不是数组，而是slice
	arr[0] = 1000
}

//此处arr为数组的地址，此时modifyArrayByPoint的空间存放的是一个地址，此地址指向TestArray01中array1空间。因此能够修改成功。
//效率高(拷地址)
func modifyArrayByPoint(arr *[2]int) {
	(*arr)[0] = 689
	//arr[0] = 689
}
