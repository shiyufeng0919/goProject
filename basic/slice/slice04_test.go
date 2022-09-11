package slice

import (
	"fmt"
	"testing"
)

/*
切片练习
*/
func TestSlice04(t *testing.T){
	/*
	编写一个函数fbn(n int),要求完成：
	1，可以接收一个n int
	2, 能够将斐波那契的数列放到切片中
	3, 提示，斐波那契的数列形式:
	arr[0] = 1  arr[1] = 1  arr[2] =2    arr[3] = 3     arr[4] = 5    arr[5] = 8 (前两个元素相加)
	*/
	slice41()
}

func slice41(){
	fbn(5)
}

func fbn(n int){
  var slice = make([]uint64,n)
  for i:=0 ; i < n ; i++ {
  	if i == 0 || i == 1{
  		slice[i] = 1
	}else{
		slice[i] = slice[i-1] + slice[i-2]
	}
  }
  fmt.Printf("slice41=== slice:%+v",slice) //[1 1 2 3 5]
}