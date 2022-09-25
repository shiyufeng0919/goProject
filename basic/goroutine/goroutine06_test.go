package goroutine

import (
	"fmt"
	"testing"
	"time"
)

/*
应用实例：
统计1-10000的数字中，哪些是素数？
1，传统方法: 循环判断各个数是否为素数
2，并发/并行方式，将统计素数任务分配给多个goroutine去完成。提升效率

小结：
使用go协程后，执行的速度，至少提升4倍。(关键看CPU个数)

注意：
如:cpu本身4个，开启4个协程则分配到了4个cpu,再多开协程也只是作用在4个cpu上，速度提升不会有明显变化。
*/
func TestGoroutine06(t *testing.T) {
	goroutine61()
}

func goroutine61() {
	intChan := make(chan int, 5000)
	resultChan := make(chan int, 2000) //放入结果
	exitChan := make(chan bool, 4)

	//任务开始时间
	startTime := time.Now().UnixNano()

	go putData(intChan)
	//开启4个goroutine,从intChan取出数据并判断是否为素数
	//并行方式，计算素数
	for i := 1; i <= 4; i++ {
		go getData(intChan, resultChan, exitChan)
	}
	//从channel中取4个结果，即可关闭
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		//任务结束时间，即:拿从exitChan拿到4个数据即算完成.(打印因耗时，可去掉)
		endTime := time.Now().UnixNano()
		fmt.Println("使用协程耗时:", endTime-startTime)
		//当成功从exitChan取出4个结果，则可以放心的关闭resultChan
		close(resultChan)
	}()
	//取出结果
	for {
		_, ok := <-resultChan
		if !ok {
			break
		}
		//fmt.Printf("素数=%d\n", res)
	}
	fmt.Println("main线程退出")
}

//channel是引用数据类型
func putData(intChan chan int) {
	for i := 1; i <= 1000; i++ {
		intChan <- i
	}
	close(intChan)
}

//素数:
func getData(intChan chan int, resultChan chan int, exitChan chan bool) {
	var flag bool //标识是否为素数
	for {
		num, ok := <-intChan
		//取不到了
		if !ok {
			break
		}
		flag = true
		//判断num是否为素数(素数:指的是“大于1的整数中，只能被1和这个数本身整除的数)
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			resultChan <- num
		}
	}
	//close(resultChan) //注意：不能关闭resultChan，因为其它协程可能还在处理
	exitChan <- true
	//close(exitChan)   //注意：此处也不能关闭
}
