package flag

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

/*
命令行参数
os.Args (http://doc.golang.ltd-> os-> Variables-> var Args []string Args保管了命令行参数，第一个是程序名。)
os.Args是一个string的切片，用来存储所有的命令行参数

flag包用来解析命令行参数.(参数顺序可以随意)
	func IntVar(p *int, name string, value int, usage string)
	IntVar用指定的名称、默认值、使用信息注册一个int类型flag，并将flag的值保存到p指向的变量。
	func StringVar(p *string, name string, value string, usage string)
	StringVar用指定的名称、默认值、使用信息注册一个string类型flag，并将flag的值保存到p指向的变量。
*/

func TestFlag01(t *testing.T) {
	//os.Args基本使用(原生,参数必须按顺序一一指定)
	flag11()
	//flag包用来解析命令行参数.(参数顺序可以随意)
	flag12()
}

func flag11() {
	fmt.Println("flag11===\n命令行参数有:", len(os.Args))
	//遍历os.Args切片，就可以得到所有的命令行输入参数值
	for k, v := range os.Args {
		fmt.Printf("args[%v]=%v \n", k, v)
	}
}

//eg: mysql -u root -pwd 123456 -h localhost -P 3306
func flag12() {
	//定义几个变量，用于接收命令行的参数值
	var user, password, host string
	var port int
	/*
		p -> *string : &user 接收用户命令行中输入的-u后面的参烽值
		name : -u 指定参数
		value: 默认值
		usage: 说明
	*/
	flag.StringVar(&user, "u", "", "用户名,默认为空")
	flag.StringVar(&password, "pwd", "", "密码,默认为空")
	flag.StringVar(&host, "h", "localhost", "主机,默认为localhost")
	flag.IntVar(&port, "P", 3306, "端口号，默认为3306")
	//这里是一个非常重要的操作，转换，必须使用该方法.
	flag.Parse() //从os.Args[1:]中解析注册的flag。必须在所有flag都注册好而未访问其值时执行。未注册却使用flag -help时，会返回ErrHelp。
	fmt.Printf("user:%s,password:%s,host:%s,port:%v", user, password, host, port)
}
