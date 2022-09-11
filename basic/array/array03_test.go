package array

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/*
 数组的复杂使用
*/

func TestArray03(t *testing.T) {
	//1.随机生成5个数，并将其反转打印. 以下为4种实现方式.
	array31()
	fmt.Println("=====")
	array32()
	fmt.Println("=====")
	array33()
	fmt.Println("=====")
	array34() //针对array33()进行优化，提升效率!
}

func array31(){
	var array [5]int
	for i:=0 ; i< 5 ; i++ {
		array[i] = rand.Intn(100) //随机生成[0-100)以内的数
	}
	fmt.Println("随机生成5个数:",array) //随机生成5个数: [81 87 47 59 81].问题点:发现每次生成都是一样的。
	for j:=len(array) -1; j >= 0; j-- {
		fmt.Println(array[j])
	}
}

func array32(){
	var array [5]int
	//为了每次生成的随机数不一样，需要给一个seed值
	rand.Seed(time.Now().UnixNano()) //利用时间(每次肯定不一样)作为seed
	for i:=0 ; i< 5 ; i++ {
		array[i] = rand.Intn(100) //随机生成[0-100)以内的数
	}
	fmt.Println("随机生成5个数:",array)
	for j:=len(array) -1; j >= 0; j-- {
		fmt.Println(array[j])
	}
}

func array33(){
	var array [5]int
	//为了每次生成的随机数不一样，需要给一个seed值
	rand.Seed(time.Now().UnixNano()) //利用时间(每次肯定不一样)作为seed
	for i:=0 ; i< 5 ; i++ {
		array[i] = rand.Intn(100) //随机生成[0-100)以内的数
	}
	fmt.Println("随机生成5个数:",array)
	temp := 0
	//交换
	for i:=0 ; i< len(array) /2 ; i++ { //len(array)为内嵌函数，会耗费资源，每调一个function，都会开辟一个独立的空间。因此此处需要优化。见array34()
		temp = array[len(array) - 1 - i]
		array[len(array) - 1 - i] = array[i]
		array[i] = temp
	}
	fmt.Println(array)
}

func array34(){
	var array [5]int
	//为了每次生成的随机数不一样，需要给一个seed值
	rand.Seed(time.Now().UnixNano()) //利用时间(每次肯定不一样)作为seed
	for i:=0 ; i< 5 ; i++ {
		array[i] = rand.Intn(100) //随机生成[0-100)以内的数
	}
	fmt.Println("随机生成5个数:",array)
	arrayLen := len(array) //效率会更高
	temp := 0
	//交换
	for i:=0 ; i< arrayLen /2 ; i++ {
		temp = array[arrayLen- 1 - i]
		array[arrayLen - 1 - i] = array[i]
		array[i] = temp
	}
	fmt.Println(array)
}
