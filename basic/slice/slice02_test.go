package slice

import (
	"fmt"
	"testing"
)

/*
 1,切片的3种使用方式.
 2,方式1和方式2区别:
   方式1是直接引用数组，这个数组是事先存在的，程序员是可见的。
   方式2是通过make来创建切片，make也会创建一个数组，是由切片在底层进行维护，程序员是不可见的。
 3,切片遍历-2种形式(for 和 for-range)
 4,切片使用注意事项及细节
	(1).切片初始化时，var slice = array[startIndex:endIndex]。从array数组下标为startIndex取到下标endIndex的元素,但不饮食array[endIndex]
    (2).切片初始化时，不能越界，范围在[0-len(array)]之间，但可动态增长。
	(3).var slice = array[0:end] 可简写 var slice = array[:end]
 	(4).var slice = array[start:len(array)] 可简写 var slice = array[start:]
	(5).cap是一个内置函数，用于统计切片容量，即最大可以放多少元素
	(6).切片定义完后,还不能使用，因为本身是一个空的。需要让其引用一个数组或make一个空间供切片来使用
	(7).切片可以继续切片
	(8).用append内置函数，可以对切片进行动态追加
		切片append的底层原理分析:
			1,切片append操作的本质就是对数组扩容
			2,go底层会创建一个新的数组newArr
			3,将slice原来包含的元素拷贝到新的数组newArr
			4,slice重新引用到newArr
			5,注意newArr是由底层维护的，程序员不可见.
    (9).切片的拷贝操作
	切片使用copy内置函数拷贝。注意必须是切片才可以拷贝，数组不可以！！！
	(10).切片是引用类型，所以在传递时，遵守引用传递机制。
*/

func TestSlice02(t *testing.T){
	//1,定义一个切片，然后让切片去引用一个已经创建好的数组。
	slice21()

	//2.通过make来创建切片
	slice22()

	//3.定义一个切片，直接就指定具体数组，使用原理类似make方式
	slice23()

	//4,切片遍历
	slice24()

	//5.append动态扩容及append底层原理分析
	slice25()

	//6.切片拷贝操作
	slice26()

	//7.示例，切片是引用数据类型
	slice27()
}

func slice21(){
	var array [5]int = [5]int{1,2,3,4,5}
	slice := array[1:3]
	fmt.Printf("\nslice21=== array:%v,slice:%v",array,slice) //array:[1 2 3 4 5],slice:[2 3]
}

func slice22(){
	/*
	1,定义: var 切片名 []type = make([]type,len,[cap]) //cap为可选项,cap >= len
	2,make为内置函数.在builtin包下[http://doc.golang.ltd]
	3,对于切片，必须make后才能使用，否则报错.
	4,内存图
	var slice []float64 = make([]float64,2,5) -> 开辟一个空间
	 slice指向一个空间 -> | 地址  | 长度2 ｜ 容量5
	                        |
							^  指向数组[0,0]
					   | 0  |  0 |     ->该数组是由make创建的，底层维护，程度员不可见。访问make创建的数组，只能通过切片来访问.
	5,通过make方式创建切片可以指定切片的大小和容量
	6,若没有给切片各无素赋值，则会使用默认零值. [int/float -> 0, string -> "", bool -> false]
	7,通过make方式创建的节片对应的数组由make底层维护，对外不可见。即只通过slice去访问各个元素.
	*/
	var slice []float64 = make([]float64,2,5) //长度=2,容量=5
	slice[1] = 88
	fmt.Printf("\nslice22=== slice:%v,len:%v,cap:%v",slice, len(slice), cap(slice)) //slice:[0 88],len:2,cap:5
}

func slice23(){
	/*
	 定义一个切片，直接指定具体数组。因[]无长度，因此不是数组，而是一个slice.
	*/
	var slice []int = []int{1,3,4}
	fmt.Printf("\nslice23=== slice:%v,len:%v,cap:%v",slice, len(slice), cap(slice)) //slice:[1 3 4],len:3,cap:3
}

func slice24(){
	var slice []float64
	slice = make([]float64,3,5)
	slice[0] = 11.0
	slice[1] = 22.0
	slice[2] = 33.0
	//for循环
	for i:=0 ; i< len(slice); i++ {
		fmt.Println("\nslice24===slice:",slice[i])
	}

	fmt.Println("")
	//切片可以继续切片
	var slice2 []float64 = make([]float64,2)
	slice2 = slice[:2]
	slice2[0] = 888
	fmt.Println("\nslice24=== modify slice:",slice) //modify slice: [888 22 33]
	//for-range
	for k,v := range slice2 {
		fmt.Printf("\nslice24=== k:%v,v:%v \n",k,v)
	}
}

func slice25(){
	/*
	切片append操作的底层原理分析:
	slice引用一个数组
	slice   | 数组array首元素地址值| 3(len) | 3(cap) |
				|
				^
	array   |  10  |   20  |  30  |
	*/
	var slice []int = []int{10,20,30}
	/*
	slice   | 数组array首元素地址值| 3(len) | 3(cap) |
					|
					^
		array   |  10  |   20  |  30  |

		执行: slice = append(slice, 40,50,60),创建一个新的数组newArray1,该数组有6个元素

			newArray1  |  0  |   0  |  0  |   0  |   0  |  0  |

			slice将原数组10-30及新追加元素40-60向newArray1拷贝，形成newArray2. 并将新数组newArray2赋值给slice.则slice的首地址将重新指向新数组.

			newArray2  |  10  |   20  |  30  |   40  |   50  |  60  |

	   执行: slice2 := append(slice, 70),在原slice 6个元素基础上，扩容一下元素.并赋初值0。将slice数组内容拷贝到newArray3，形成newArray4.
			此时，赋值给了slice2，则slice2指向了newArray4首地址，而slice还是保留原newArray2首地址不变。因此slice2为newArray4值，而slice2
		    为新newArray4值.

			newArray3  |  0  |   0  |  0  |   0  |   0  |  0  |  0  |

			newArray4   |  10  |   20  |  30  |   40  |   50  |  60  |   70  |
	*/
	//1,切片追加元素
	slice = append(slice, 40,50,60)
	fmt.Printf("\nslice25===slice:%+v",slice) //slice25===slice:[10 20 30 40 50 60]

	slice2 := append(slice, 70)
	fmt.Printf("\nslice25===slice:%+v,slice2:%+v",slice,slice2) //slice25===slice:[10 20 30 40 50 60],slice2:[10 20 30 40 50 60 70]

	//2,切片追加切片
	var slice3 []int = []int{88,99}
	/*
	注意点： 1,切片仅能追加切片，不能追加数组. 2,切片追加切片...规定用法。
	*/
	slice = append(slice, slice3...)
	fmt.Printf("\nslice25===append slice:%+v",slice) //slice25===append slice:[10 20 30 40 50 60 88 99]

}

func slice26(){
	/*
	1,创建数组slice1并赋值 | 888  |  999 |
	*/
	var slice1 []int = []int{888,999}
	/*
	2,创建数组slice2并赋值 | 0  |  0 |  0  |  0 |
	*/
	var slice2 = make([]int,4)
	/*
	3,将组数slice1拷贝给slice2,从第一个元素开始拷贝。一直结束，未填充的保留原值. | 888  |  999 |  0  |  0 |
	*/
	copy(slice2,slice1)
	/*
	上述说明:
	1,copy(param1,param2) 参数的数据类型是切片
	2,slice1和slice2的数据空间是独立的，互不影响，即slice1[0] = 99不会影响slice2[0]
	*/
	fmt.Printf("\n slice26=== slice1:%+v",slice1) //[888 999]
	fmt.Printf("\n slice26=== slice2:%+v",slice2) //[888,999,0,0]

	var slice3 = make([]int,1)
	copy(slice3,slice1) //将slice1拷贝给slice3,仅拷贝1个元素即可，第二个不需要了，没有问题.
	fmt.Printf("\n slice26=== slice3:%+v",slice3) //[888]
}

func slice27(){
	/*
	定义一个切片，此时slice还不能使用，该切片没有空间
	*/
	var slice []int
	/*
	定义一个数组array  [1,2,3,4,5]
	*/
	var array [5]int = [...]int{1,2,3,4,5}
	/*
	让slice指向数组array
	*/
	slice = array[:]
	/*
	将slice交给slice2,此时slice和slice2均指向array数组
	*/
	var slice2 = slice
	/*
	通过slice2将第一个元素改变，由于切片slice是引用数据类型，所以改变了原array的第一个元素值，而slice和slice2又同时指向array数组，因此第一个元素均发生变化。
	*/
	slice2[0] = 888
	fmt.Printf("\nslice27===slice:%+v",slice) //[888 2 3 4 5]
	fmt.Printf("\nslice27===slice2:%+v",slice2) //[888 2 3 4 5]
	fmt.Printf("\nslice27===array:%+v",array)   //[888 2 3 4 5]

	var slice3 = []int{1,2,3,4}
	fmt.Printf("\nslice27===slice3:%+v",slice3) //[1 2 3 4]
	modifySlice(slice3) //slice3实参
	fmt.Printf("\nslice27===modify slice3:%+v",slice3) //999 2 3 4]
}

//将切片传给一个函数，该切片是引用数据类型
func modifySlice(slice []int){ //slice形参
	/*
	这里slice改变会改变实参,即函数内改变会影响函数外.
	*/
	slice[0] = 999
}