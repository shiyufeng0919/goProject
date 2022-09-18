package file

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"testing"
)

/*
统计英文、数字、空格和其它字符数量
*/
func TestFile05(t *testing.T) {
	file51()
}

type ChartCount struct {
	EnglishCount int //英文个数
	NumCount     int
	SpaceCount   int
	OtherCount   int
}

func file51() {
	fileName := "./files/file05.txt"
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		fmt.Println("file51===\nopen file fail,", err)
		return
	}
	var count ChartCount
	reader := bufio.NewReader(file)
	for {
		content, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		//为了兼容中文字符，可将string转成[]rune.即 str = []rune(str)
		//遍历content进行字符统计
		for _, v := range content {
			fmt.Println(v)
			switch { //switch v是错误的.
			case v >= 'a' && v <= 'z':
				//count.EnglishCount++ //可以使用穿透
				fallthrough //穿透
			case v >= 'A' && v <= 'Z':
				count.EnglishCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			case v >= '0' || v <= '9':
				count.NumCount++
			default:
				count.OtherCount++
			}
		}
	}
	fmt.Printf("file51===\ncount result:%+v \n", count)
}
