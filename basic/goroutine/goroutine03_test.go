package goroutine

import (
	"fmt"
	"testing"
)

/*
channel(管道) 「channel是引用数据类型」

channel最大价值就是一边往里放，一边往出取，类似于水管一样.

一，channel介绍
1,channel本质上就是一个数据结构-队列 (可以用链表/map/数组等来实现队列)
2,数据是先进先出(与栈相反，栈是先进后出)
3,线程安全，多goroutine访问时，不需要加锁，即channel本身就是线程安全的.(即多个协程在操作同一channel时不会发生资源竞争问题)
4,channel是有类型的，一个string的channel只能存放string类型数据

二，channel基本使用
1,定义/声明
var 变量名 chan 数据类型
eg:
var intChan chan int    //intChan用于存放int类型数据
var mapChan chan map[int]string   //mapChan用于存放map[int]string类型数据
var perChan chan Person      //perChan用于存放Person结构体
var perChan2 chan *Person   //perChan2用于存放Person结构体指针
...
说明：
1),channel是引用数据类型
2),channel必须初始化才能写入数据，即make后才能使用
3),管道是有类型的，如: intChan只能写入int类型数据

2,channel初始化
使用make进行初始化
var intChan chan int
intChan = make(chan int,10) //channel中最多存放10个int类型数据

小结：
1,channel中只能存放指定的数据类型
2,channel的数据存放满后，就不能再放入了
3,如果从channel取出数据后，可以继续放入
4,在没有协程的情况下，若channel数据取完了，再取，就会报Fatal error: all goroutines are asleep - deadlock!
*/

func TestGoroutine03(t *testing.T) {
	//向channel写入数据和从channel读取数据（基本数据类型）
	goroutine31()
	//向channel写入数据和从channel读取数据（map数据类型）
	goroutine32()
	//向channel写入数据和从channel读取数据（struct数据类型）
	goroutine33()
	//向channel写入数据和从channel读取数据（struct类型指针）
	goroutine34()
	//向channel写入数据和从channel读取数据（任意数据类型变量)-使用空接口类型，再操作struct时，需要应用类型断言来取出struct字段）
	goroutine35()
}

func goroutine31() {
	/*
		说明: channel是引用数据类型
						0xc00010e028(本身地址)
		intChan  ---->  0xc000132100(intChan指向地址)
	*/
	var intChan chan int
	intChan = make(chan int, 3)
	fmt.Printf("intChan===\n intChan的值:%+v,intChan本身地址:%p\n", intChan, &intChan) // intChan的值:0xc000132100,intChan本身地址:0xc00010e028

	/*
		注意:当给channel写入数据时，不能超过其容量.若非要放置超过容量数据，则需要先从channel取出数据(长度减少，但容量不变)，再向里放入数据。
	*/
	intChan <- 99
	num := 88
	intChan <- num
	fmt.Printf("intChan len:%v,cap:%v\n", len(intChan), cap(intChan)) //intChan len:2,cap:3

	var num1 int
	num1 = <-intChan
	fmt.Println("num1:", num1)                                               //num1: 99
	fmt.Printf("取出数据1后 intChan len:%v,cap:%v\n", len(intChan), cap(intChan)) //取出数据后 intChan len:1,cap:3

	//取出第2个数据，但是不接收。扔掉了。
	<-intChan

	//在「没有协程」的情况下，若channel数据已全部取出，再取就会报错
	//var num3 int
	//num3 = <-intChan
	//fmt.Println("取一个没有的数据num3:", num3) //fatal error: all goroutines are asleep - deadlock!
}

func goroutine32() {
	var mapChan chan map[int]string
	mapChan = make(chan map[int]string, 10)
	map1 := make(map[int]string)
	map1[1] = "hi"
	map1[2] = "syf"
	map2 := make(map[int]string)
	map2[1] = "hello"
	map2[2] = "channel"
	mapChan <- map1
	mapChan <- map2
	fmt.Println("goroutine32===\nmapChan:\n", mapChan)                        //mapChan: 0xc00007a240
	fmt.Printf("mapChan.len:%v,mapChan.cap:%v\n", len(mapChan), cap(mapChan)) //mapChan.len:2,mapChan.cap:10

	map11 := <-mapChan
	map12 := <-mapChan
	fmt.Printf("取出第一个数据map11:%+v,取出第二个数据map12:%+v", map11, map12)
}

type Person struct {
	Name string
	Age  int
}

func goroutine33() {
	var structChan chan Person
	structChan = make(chan Person, 10)
	p1 := Person{Name: "syf", Age: 18}
	structChan <- p1
	fmt.Printf("goroutine33===\nstructChan:%v,structChan.len:%v,structChan.cap:%v \n", structChan, len(structChan), cap(structChan)) //structChan:0xc0000ae360,structChan.len:1,structChan.cap:10
	p2 := Person{Name: "channel", Age: 19}
	structChan <- p2
	fmt.Printf("p2->structChan:%v,structChan.len:%v,structChan.cap:%v \n", structChan, len(structChan), cap(structChan)) //p2->structChan:0xc0000ae360,structChan.len:2,structChan.cap:10
	data := <-structChan
	fmt.Printf("从structChan取出数据:%+v \n", data) //从structChan取出数据:{Name:syf Age:18}
}

func goroutine34() {
	var structChan chan *Person
	structChan = make(chan *Person, 10)
	p1 := Person{Name: "syf", Age: 18}
	structChan <- &p1
	fmt.Printf("goroutine34===\np1->structChan:%v,structChan.len:%v,structChan.cap:%v \n", structChan, len(structChan), cap(structChan)) //p1->structChan:0xc000056420,structChan.len:1,structChan.cap:10
	p2 := Person{Name: "channel", Age: 19}
	structChan <- &p2
	fmt.Printf("p2->structChan:%v,structChan.len:%v,structChan.cap:%v \n", structChan, len(structChan), cap(structChan)) //p2->structChan:0xc000056420,structChan.len:2,structChan.cap:10
	data := <-structChan
	fmt.Printf("从structChan取出数据:%+v \n", data) //从structChan取出数据:&{Name:syf Age:18}
}

func goroutine35() {
	var allChan chan interface{}
	allChan = make(chan interface{}, 10)
	data1 := 88
	data2 := "this is string"
	data3 := make(map[int]string)
	data3[1] = "this is"
	data3[2] = "map"
	data4 := Person{Name: "struct", Age: 28}
	allChan <- data1
	allChan <- data2
	allChan <- data3
	allChan <- data4
	fmt.Printf("goroutine35===\nallChan:%v,len:%v,cap:%v\n", allChan, len(allChan), cap(allChan))

	data11 := <-allChan
	data22 := <-allChan
	data33 := <-allChan
	data44 := <-allChan
	fmt.Println("data11:", data11) //data11: 88
	fmt.Println("data22:", data22) //data22: this is string
	fmt.Println("data33:", data33) //data33: map[1:this is 2:map]
	fmt.Println("data44:", data44) //data44: {struct 28}
	//取出data44结构体的名称字段 {不能data44.Name这样用，编译器通不过}
	name := data44.(Person).Name      //应用类型断言，因为data44是一个interface{}空接口类型，需要应用类型断言转换成具体的类型。
	fmt.Println("data44 name:", name) //data44: {struct 28}
}
