package array

import (
	"fmt"
	"testing"
)

/*
数组应用案例:
*/
func TestArray02(t *testing.T) {
	//1.创建一个byte类型的26个元素的数组，分别放置'A'-'Z'.使用for循环访问所有元素并打印出来。提示：字符数据运算'A'+1 -> 'B'
	array21()

	fmt.Println()

	//2.请求出一个数组的最大值，并得到对应的下标
	array22()

	fmt.Println()

	//3.请求出一个数组的和及平均值. for-range
	array23()
}

func array21(){
	var array [26]byte
	for i:=0 ; i< 26; i++ {
		array[i] = 'A' + byte(i)  //注意此处，i需要转换成byte类型，否则会报错.
	}
	fmt.Printf("array:%v \n",array)

	for _,v := range array {
		fmt.Printf("%c",v)
	}
}

func array22(){
	var array [3]int = [3]int{88,99,66}
	var maxValue int
	var maxIndex int
	for k,v := range array {
		if v > maxValue {
			maxValue = v
			maxIndex = k
		}
	}
	fmt.Printf("maxIndex:%v,maxValue:%v",maxIndex,maxValue)
}

func array23(){
	var array [3]int = [...]int{1,2,4}
	sum := 0
	for _,v := range array {
		sum += v
	}
	fmt.Printf("数组的和:%v,平均值:%v",sum,sum/len(array)) //7,2 (如何让平均值保留到小数点)

	fmt.Printf("数组的和:%v,平均值:%v",sum,float64(sum)/float64(len(array))) //将int -> float64 (golang的数据类型为显示转换)
}
