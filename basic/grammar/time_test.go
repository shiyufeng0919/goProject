package grammar

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	//13位时间戳(1663579789807)转time.Time
	time01()
}

func time01() {
	/*
		10位数的时间戳是以秒为单位，如: 1663579789
		13位数的时间戳是以毫秒为单位，如: 1663579789807
		19位数的时间戳是以纳秒为单位，如: 1664098640986890000
	*/
	//获取当前时间或时间戳
	unix := time.Now().Unix()
	unixNano := time.Now().UnixNano()
	fmt.Printf("unix:%v,unixNano:%v\n", unix, unixNano) //unix:1664098772,unixNano:1664098772468374000

	//13位时间戳转时间格式
	dataTime, _ := UnixToTime("1663579789807")
	fmt.Println("dataTime:", dataTime) //dataTime: 2022-09-19 17:29:49 +0800 CST
	formatDataTime := dataTime.Format("2006-01-02 15:04:05")
	fmt.Println("formatDataTime:", formatDataTime) //formatDataTime: 2022-09-19 17:29:49

	unixData := TimeToUnix(dataTime) //1663608589000
	fmt.Println("unixData:", unixData)

	fmt.Println("-----time01 end-----")
}
func UnixToTime(e string) (dataTime time.Time, err error) {
	data, err := strconv.ParseInt(e, 10, 64)
	dataTime = time.Unix(data/1000, 0)
	return
}

func TimeToUnix(e time.Time) int64 {
	timeUnix, _ := time.Parse("2006-01-02 15:04:05", e.Format("2006-01-02 15:04:05"))
	return timeUnix.UnixNano() / 1e6
}
