package sort

import (
	"fmt"
	"testing"
)

/*
查找:
1，顺序查找
2，二分查找(效率高，利用到了递规.在一个有序数组中进行二分查找，若数组不是有序的，则不能使用二分查找)

案例：
1,有一个数列: 张三，李四，王五，马六
猜数游戏，从键盘中任意输入一个名称，判断数列中是不哭包含此名称[顺序查找]

2,请对一个有序数组进行二分查找{1,8,10,89,1000,1234}。输入一个数，看看该数组是否存在此数，
且求出下标，若没有就提示"no data".[会使用到递归]
*/
func TestSort02(t *testing.T){
	//案例1,顺序查找
	sort21()

	//案例2,二分查找
	sort22()
}

func sort21(){
	var names [4]string = [4]string{"张三","李四","王五","马六"}
	var goalName = "马六"
	//顺序查找1
	for k,v := range names {
		if v == goalName {
			fmt.Printf("\nsort21===下标:%v,值:%s",k,v)
			break
		}else if k == len(names) - 1 {
			fmt.Printf("\nsort21===没有找到:%s",goalName)
		}
	}

	//顺序查找2(推荐)
	index := -1
	for k,v := range names {
		if v == goalName {
		   index = k
		   break
		}
	}
	if index == -1 {
		fmt.Printf("\nsort21===没有找到:%s",goalName)
	}else{
		fmt.Printf("\nsort21===下标:%v,值:%s",index,goalName)
	}
}

func sort22(){
	/*
	二分查找，数组应该是一个有序的，无序的不行.
	思路:
	1,array是一个有序数组，从小->大
	2,找到中间下标 midIndex = (leftIndex + rightIndex)/2,将该下标值与findVal比较
	3,array[midIndex] > findVal,则向leftIndex ~ midIndex-1 间查找
	  array[midIndex] < findVal,则向midIndex+1 ~ rightIndex 间查找
	4,什么情况下，说明找不到(分析退出递归条件)
	  leftIndex > rightIndex 情况下，需要退出递归.
	*/
	var array [6]int = [...]int{1,8,10,89,1000,1234}
	var findVal int = 89
	binaryFind(&array,0,len(array) - 1 ,findVal)
}

//二分查找
func binaryFind(arr *[6]int,leftIndex,rightIndex,findVal int){
	if leftIndex > rightIndex {
		fmt.Printf("\nbinaryFind=== 找不到")
		return
	}
	//先找到中间下标
	midIndex := (leftIndex + rightIndex)/2
	if (*arr)[midIndex] > findVal {
		//leftIndex ~ midIndex -1 间查找
		binaryFind(arr,leftIndex,midIndex-1,findVal)
	}else if (*arr)[midIndex] < findVal{
		// midIndex + 1 ~ rightIndex 间查找
		binaryFind(arr,midIndex+1,rightIndex,findVal)
	}else {
		fmt.Printf("\n binaryFind=== 找到了,下标为:%v",midIndex)
	}
}

