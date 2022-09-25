package goroutine

import (
	"fmt"
	"testing"
)

/*
channel使用细节和注意事项：
1,channel可以声明为只读，或者只写性质。
应用场景：可以定义一个只写协程专门写(putData)和一个只读协程专门读(getData)
2,使用select可以解决从管道取数据的阻塞问题
3,goroutine中使用recover,解决协程中出现panic,导致程序崩溃问题
//在可能导致panic前执行defer recover捕获panic.这样不会影响主线程的执行。
defer func() {
		if err := recover(); err != nil {
			fmt.Println("发生错误:", err)
		}
}()
*/
func TestGoroutine07(t *testing.T) {
	//channel可以声明为只读，或者只写性质
	goroutine71()
	//使用select可以解决从管道取数据的阻塞问题
	goroutine72()
}

func goroutine71() {
	//默认情况下，管道是双向的
	var chan1 chan int //可读可写
	chan1 = make(chan int, 10)
	chan1 <- 100
	<-chan1

	var writeChan chan<- int       //只写
	writeChan = make(chan int, 10) //注意：类型依然是chan int,只是定义的类型有区分
	writeChan <- 99
	//data :=<-writeChan //编译错误，「chan<-」只能写，不能读

	var readChan <-chan int //只读
	readChan = make(chan int, 10)
	<-readChan //此处会发生fatal error: all goroutines are asleep - deadlock! 可使用defer recover解决
}

func goroutine72() {
	var intChan chan int
	intChan = make(chan int, 10)
	for i := 1; i <= 10; i++ {
		intChan <- i
	}
	var strChan chan string
	strChan = make(chan string, 5)
	for i := 1; i <= 5; i++ {
		strChan <- fmt.Sprintf("syf%d", i)
	}
	/*
		传统方法在遍历管道时，若不关闭会阻塞而导致deadlock.
		有时在不好确定什么时候关闭管道时，该怎么办呢？可以使用select方式解决!!!
	*/
	//label:
	for {
		select {
		case v := <-intChan: //注意:这里如果intChan一直没有关闭，也不会一直阻塞，它会向下一个case匹配
			fmt.Println("从intChan读取数据:", v)
		case v := <-strChan:
			fmt.Println("从strChan读取数据:", v)
		default:
			fmt.Println("都取不到，可写自己的逻辑")
			//break //此处直接break只能退出select，并不能退出for
			//break label //break到指定标签，不建议使用
			return //可直接返回,下面代码没有机会再执行
		}
	}
}
