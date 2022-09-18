package file

import (
	"fmt"
	"os"
	"testing"
)

/*
判断文件或目录是否存在使用os.Stat()函数返回的错误值进行判断
1,若返回的错误为nil,则文件或文件夹存在
2,若返回的错误类型使用os.IsNotExist()判断为true，则文件或文件夹不存在
3,若返回的错误为其它类型，则不确定是否存在
*/

func TestFile03(t *testing.T) {
	//判断文件/目录是否存在
	file31()
}

func file31() {
	fileName := "./files/file03_test.txt"
	exists, err := PathExists(fileName)
	if err != nil {
		fmt.Println("file31===judge file is exists fail,", err)
		return
	}
	fmt.Printf("file31===judge file is exists:%v \n", exists)

	dirName := "./files"
	exist, err := PathExists(dirName)
	if err != nil {
		fmt.Println("file31===judge dir is exists fail,", err)
		return
	}
	fmt.Printf("file31===judge dir is exist:%v \n", exist)
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { //文件或目录存在
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
