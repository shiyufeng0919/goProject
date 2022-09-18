package file

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

/*
一，文件
1,文件最主要作用就是保存数据，即文件是数据源.
2,文件在程序中是以流的形式来操作的。

Go程序    <-------输入流(读文件)-----    文件
(内存)	 --------输出流(写文件)----->
1）流：数据在数据源(文件)和程序(内存)之间经历的路径。
2）输入流：数据从数据源(文件)到程序(内存)的路径。
3）输出流：数据从程序(内存)到数据源(文件)的路径。

二，读文件

1,os.File封装所有文件相关操作，File是一个结构体
http://doc.golang.ltd -> os -> File-> 找到File结构体
	type File
		func Create(name string) (file *File, err error)
		func Open(name string) (file *File, err error) #Open打开一个文件用于读取
		.....
		func (f *File) Close() error #文件打开后，一定要close
		.....

http://doc.golang.ltd -> package bufio
bufio包实现了有缓冲的I/O。它包装一个io.Reader或io.Writer接口对象，创建另一个也实现了该接口，且同时还提供了缓冲和一些文本I/O的帮助函数的对象。
	type Reader
		...
	type Writer
		...
*/

func TestFile01(t *testing.T) {
	//打开文件
	file11()
	//读文件(带缓冲):读取一部分文件内容到内存，处理完成后再继读取,直到文件末尾。(适合大文件)
	file12()
	//读文件: 一次性读取文件到内存.(不适合大文件)
	file13()
}

func file11() {
	//打开文件
	fileName := "./files/file01.txt"
	//file: 叫file对象或file句柄或file指针
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		fmt.Println("open file fail,", err)
		return
	}
	fmt.Printf("file11===\nfile=%v \n", file) //file=&{0xc0000562a0}
}

/*
此示例有个问题：文件中最后一行汉字部分未打印.
*/
func file12() {
	//打开文件
	fileName := "./files/file01.txt"
	//file: 叫file对象或file句柄或file指针
	file, err := os.Open(fileName)
	//在函数退出之前，会调用defer后面的代码。此处即时关闭file句柄，否则会有内存泄漏.
	defer file.Close()
	if err != nil {
		fmt.Println("file12===\nopen file fail,", err)
		return
	}
	fmt.Println("file12===")
	//创建一个 *Reader，是带缓冲的.默认缓冲区为4096。好处为:在读取文件不是一次性读取到内存，而是读一部分处理一部分。
	reader := bufio.NewReader(file)
	//循环读取文件的内容
	for {
		content, err := reader.ReadString('\n') //读取到一个换行就end.
		if err == io.EOF {                      //io.EOF代表读取到文件末尾
			break
		}
		fmt.Print(content) //注意此处不要用fmt.Println，因为reader.ReadString('\n')已加一个换行符，若用fmt.Println，则又会添加一个换行符.
	}
	fmt.Println("file read finish!")
}

func file13() {
	fileName := "./files/file01.txt"
	/*
		注意：因为没有显示的open文件，所以也不需要显示的close文件。因为文件的open和close被封装到ReadFile函数内部
	*/
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("file13===\nread file fail,", err)
		return
	}
	fmt.Println("file13===\nread content:", string(content))
}
