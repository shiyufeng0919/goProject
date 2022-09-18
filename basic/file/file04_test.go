package file

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"testing"
)

/*
拷贝文件
*/

func TestFile04(t *testing.T) {
	//拷贝图片到另一目录 | io包
	file41()
}

func file41() {
	srcFileName := "./files/file04.JPG"
	dstFileName := "./files/img/merry.jpg"
	if _, err := copyFile(dstFileName, srcFileName); err != nil {
		fmt.Println("copy fail,", err)
		return
	}
	fmt.Println("copy finish.")
}

//接收两个文件路径
func copyFile(dstFileName string, srcFileName string) (written int64, err error) {
	/*
		func Open(name string) (file *File, err error)
		Open打开一个文件用于读取。如果操作成功，[返回的文件对象的方法可用于读取数据]；对应的文件描述符具有O_RDONLY模式。如果出错，错误底层类型是*PathError。
		打开一个已经存在的文件，并对该文件进行读取!!!
	*/
	scrFile, err := os.Open(srcFileName)
	defer scrFile.Close()
	if err != nil {
		fmt.Println("open src file fail,", err)
		return 0, err
	}
	reader := bufio.NewReader(scrFile)
	/*
		func OpenFile(name string, flag int, perm FileMode) (file *File, err error)
		OpenFile是一个更一般性的文件打开函数，大多数调用者都应用Open或Create代替本函数。它会使用指定的选项（如O_RDONLY等）、指定的模式（如0666等）打开指定名称的文件。如果操作成功，返回的文件对象可用于I/O。如果出错，错误底层类型是*PathError。
		以某种模式打开一个文件，如os.O_CREATE，文件不存在则创建
	*/
	destFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	defer destFile.Close()
	if err != nil {
		fmt.Println("open file fail,", err)
		return
	}
	//通过scrFile,获取Writer
	bufio.NewWriter(destFile)
	//关键点在于拿到一个reader和一个writer.
	return io.Copy(destFile, reader)
}
