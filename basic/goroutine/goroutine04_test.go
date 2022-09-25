package goroutine

import (
	"fmt"
	"testing"
)

/*
channel的遍历和关闭
一，channel的关闭
使用内置函数close可以关闭channel,当channel关闭后，就不能再向channel写入数据，但是仍可从channel读数据。
即一旦channel关闭，则只能读不能写.

二，channel的遍历
channel支持for-range方式进行遍历，注意两个细节:
1,在遍历时，若channel没有关闭，则会出现deadlock错误
2,在遍历时，若channel已经关闭，则会正常遍历数据，遍历完后，就会退出遍历.
*/

func TestGoroutine04(t *testing.T) {
	//channel的关闭(只能读不能写)
	goroutine41()
	//channel的遍历(for-range) 注意不能使用for，理由: for i:=0;i<n;i++,而n是一直在变化的，随着数据从channel取出，长度是在一直减少。
	goroutine42()
}

func goroutine41() {
	var intChan chan int
	intChan = make(chan int, 10)
	intChan <- 100
	intChan <- 99
	intChan <- 88
	close(intChan) //用于关闭channel,内建(builtin)函数close关闭信道
	//intChan <- 66  //此时不能再写入数据到channel，因为上述已close(channel)，此处会报错.
	num1 := <-intChan                                //此时能够读数据
	fmt.Println("goroutine41===\nintChan:", intChan) //0xc0000ee000
	fmt.Println("num1:", num1)                       //num1: 100
}

func goroutine42() {
	var intChan chan int
	intChan = make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan <- i * 2
	}
	/*
			channel只返回一个值(没有下标，必须按顺序取!!!)
			注意：在遍历时，若channel没有关闭，则会出现fatal error: all goroutines are asleep - deadlock!错误
		         但是数据已全部取出了,取到最后发现没有可取的数据，因些Panic. 所以一定要关闭channel
	*/
	close(intChan)
	for v := range intChan {
		fmt.Println("goroutine42===\nv=", v)
	}
}
