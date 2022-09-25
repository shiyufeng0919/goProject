package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
channel(管道)
1,不同goroutine之间如何通讯
1)全局变量加锁同步
	package sync
	sync包提供了基本的同步基元，如互斥锁。除了Once和WaitGroup类型，大部分都是适用于低水平程序线程，高水平的同步使用channel通信更好一些。
2)channel

为什么需要channel
goroutine22()示例使用全局变量加锁同步来解决goroutine的通讯，但不完美
A,主线程在等待所有goroutine全部完成的时间很难确定，我们设置10s，仅是估算
B,若主线程休眠时间长了，会加长等待时间，若设置短了，可能还有goroutine处于工作状态，这时也会随主线程的退出而销毁
C,通过全局变量加锁同步来实现通讯，也并不利于多个协程对全局变量的读写操作
*/

func TestGoroutine02(t *testing.T) {
	//计算1-200的各个数的阶乘，并且把各个数的阶乘放入到map中，最后显示出来(会发生panic)
	//goroutine21()
	//计算1-200的各个数的阶乘，并且把各个数的阶乘放入到map中，最后显示出来
	goroutine22()
}

var (
	myMap = make(map[int]int, 10)
	//声明一个全局的互斥锁lock,Mutex:互斥
	lock sync.Mutex
)

/*
此示例会发生Panic:fatal error: concurrent map writes
go中的map不是并发安全的，所以当多个goroutine同时对map执行写操作的时候，就会报上述错误。
如何避免：
1，尽量少定义全局map变量
2，如果必须定义全局map变量，可以加锁
(1),优化1.可以采用cow策略，read不加锁，每次修改copy修改，再赋值
(2),优化2.可以采用分片锁，减小锁的粒度
*/
func goroutine21() {
	/*
			   均向一个map空间去写，产生了竞争关系->引出channel
		             -->  goroutine1 --->
		go主线程		 -->  goroutine2 --->   map空间
						     ...

		执行时，可使用go build -race main.go查看是否存在资源竞争关系!!!
	*/
	//开启多个协程来完成任务,启动200个协程对同一个map进行写操作，强制触发并发写
	for i := 1; i <= 200; i++ {
		go calculate1(i)
	}
	for k, v := range myMap {
		fmt.Printf("map[%d]=%d\n", k, v)
	}
}

func goroutine22() {
	/*
				   goroutine1拿到锁继续向下执行，完成后解锁,goroutine2过来发生有锁，则加入到队列排队，goroutine3过来也加入到队列排队...
		           goroutine1解锁后，会从队列中取出goroutine2继续执行...
			             -->  goroutine1 --->    锁
			go主线程		 -->  goroutine2 --->    map空间
							     ...			 解锁

			执行时，可使用go build -race main.go查看是否存在资源竞争关系!!!
	*/
	//开启多个协程来完成任务,启动200个协程对同一个map进行写操作，强制触发并发写
	for i := 1; i <= 200; i++ {
		go calculate2(i)
	}
	//加上休眠时间，为了避免协程还未执行完，主线程就结束了，从而导致拿到空数据.
	time.Sleep(10 * time.Second)
	lock.Lock()
	for k, v := range myMap {
		//会偶发性出现 fatal error: concurrent map iteration and map write,说明此处代码也有资源竞争问题，在for外也要加锁.
		//我们的数的阶乘很大，结果会越界,导致阶乘的输出结果为0
		fmt.Printf("map[%d]%d\n", k, v)
	}
	lock.Unlock()
}

//计算各个数的阶乘，并放入到map.
func calculate1(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	myMap[n] = res
}

//计算各个数的阶乘，并放入到map.(加锁)
func calculate2(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}
