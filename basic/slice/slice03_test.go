package slice

import (
	"fmt"
	"testing"
)

/*
string和slice
1,string底层是一个byte数组，因此string也可以进行切片处理
2,string和切片在内在的形式
3,string是不可变的，即不能通过str[0]='z'方式来修改字符串
4,若需要修改字符串，可以先string -> []byte 或 []rune -> 修改 -> 重写转成string.
*/

func TestSlice03(t *testing.T){
	//string底层是一个byte数组，因此string也可以进行切片处理
	slice31()

	//修改string处理方案(string-> []byte或[]rune{兼容中文} -> 修改 -> string)
	slice32()
}

func slice31(){
	/*
	str - >  | 0x0011  |  9(len) |
	            |
				^ 指向数组元素的首地址0x0011
			  | h | e | l | l | o | @ | s | y | f |
										^
	                					|
					--------------------指向str下标元素为6的首地址
	               |
	slice - > | 0x0059 | 3(len) |
	*/
	str := "hello@syf"
	//使用切片获取到syf,从下标为6的元素开始，切到最后。
	slice := str[6:]
	fmt.Printf("slice31=== slice:%+v",slice) //syf

	//str[0] = 888 //编译不通过，理由string是不可变的.
}

func slice32(){
	//将str下标=7的c元素修改为x,结果syf@kaixin
	str := "syf@kaicin"
	//string -> []byte
	array := []byte(str)
	array[7] = 'x'
	//[]byte -> string
	str = string(array)
	fmt.Printf("\nslice32===str:%s",str) //syf@kaixin

	/*
	上述细节：str -> []byte后，可以处理英文和数字，但是不能处理中文。
	理由: byte是按一个字节计算，而汉字为1个汉字占3个字节.因此会出现乱码
	解决方法：将string转成[]rune即可.因为rune是按字符处理的，兼容汉字.
	*/

	str2 := "kaixin@syf"
	//错误写法
	//array2 := []byte(str2)
	//array2[8] = '玉' //异常: constant 29577 overflows byte

	//正确写法
	array2 := []rune(str2)
	array2[8] = '玉'

	str2 = string(array2)
	fmt.Printf("\nslice32===str2:%s",str2) //kaixin@s玉f
}
