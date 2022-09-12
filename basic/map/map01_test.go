package _map

import (
	"fmt"
	"testing"
)

/*
「map是引用类型」

一。map是key-value数据结构，又称为字段或关联数组。类似java编程语言的集合。

二。基本语法： var map 变量名 map[keyType]valueType
key类型：eg: bool,数字,string,指针,channel,及 接口，结构体，数组。
key类型通常为int,string
注意: slice,map,function不能作为key,理由：这几个没法用==来判断

value类型: 和key一样
value类型通常为数字(整数，浮点数),string,map,struct

三。map声明示例
	var map1 map[string]string
	var map2 map[string]int
	var map3 map[int]string
	var map4 map[string]map[string]string
	//注意：「声明map是不会分配内存的」(备:数组是声明时即分配内存)，初始化需要make,分配内存后才能赋值和使用

四。map是无序的key,value。即不按key排序也不按value排序!!!

注意:
1,map在使用前一定要make
2,map的key是不能重复的，如果重复，则以最后这个key,value为主(即value被覆盖)
3,map的value是可以重复的
4,map的key-value是无序的
5,make标准库文档 http://doc.golang.ltd  -> builtin -> func make(Type, size IntegerType) Type

五。map声明的3种方式: map22()

六。map的增删改查map14() & map15()
细节说明:
1,若我们要删除map的所有key,则没有一个专门的方法一次删除，可以遍历一下key,逐个删除。
2,或者 map=make(...),make一个新的，让原来的成为垃级，被gc回收(推荐)

七。map查找
value,ok := map["key"] //查找到ok=true,否则ok=false

八。map遍历

map的遍历使用for-range.

九。map的长度

http://doc.golang.ltd  -> builtin ->  func len(v Type) int
*/
func TestMap01(t *testing.T) {
	//map的声明和注意事项
	map11()
	//map使用的方式(3种)
	map12()
	//实例: 演示一个key-value的value是map的示例，如:存放3个学生信息，每个学生有name和sex信息
	map13()
	//map的增加和更新,make["key] = value //key存在则增加，否则更新
	map14()
	//map删除某一key
	map15()
	//map删除所有key
	map16()
	//map查找
	map17()
	//map遍历 for-range
	map18()
}

func map11() {
	var map1 map[string]string
	fmt.Println("map11==,map1:", map1) //map[]
	//声明后直接赋值会panic,因为声明map是不会分配内存的，相当于还不存在的情况下就给其赋值，会报panic.
	//map1["key1"] = "value01" //panic: assignment to entry in nil map [recovered]

	//在使用map前，需要先make，make的作用就是给map分配数据空间. 如下分配3个数据空间
	map1 = make(map[string]string, 3)
	map1["key1"] = "张三"
	map1["key2"] = "李四"
	map1["key3"] = "王五"
	fmt.Println("map11===", map1) //map11=== map[key1:张三 key2:李四 key3:王五]
}

func map12() {
	//方式一: 先声明，再make，再赋值
	var map1 map[string]string
	map1 = make(map[string]string, 10)
	map1["key1"] = "syf"

	//方式二: 声明即make，再赋值。不声明空间大小，会自动增长(推荐)
	var map2 = make(map[string]string)
	map2["key2"] = "syf"
	map2["key3"] = "smile"
	map2["key4"] = "hello"

	//方式三: 声明即赋值
	//<=> map3:= map[string]string{"key5":"syf"}  //使用类型推导
	//<=> var map3 = map[string]string{"key5":"syf"}
	var map3 map[string]string = map[string]string{"key5": "syf", "key6": "smile"}
	map3["key6"] = "smile"
	fmt.Printf("\nmap22===map1:%+v,map2:%+v,map3:%+v \n", map1, map2, map3)
}

func map13() {
	var stuMap = make(map[string]map[string]string)
	stuMap["stu01"] = make(map[string]string)
	stuMap["stu01"]["name"] = "syf"
	stuMap["stu01"]["sex"] = "girl"
	stuMap["stu02"] = make(map[string]string)
	stuMap["stu02"]["name"] = "smile"
	stuMap["stu02"]["sex"] = "boy"
	fmt.Println("map13===stuMap:", stuMap)                  //map13===stuMap: map[stu01:map[name:syf sex:girl] stu02:map[name:smile sex:boy]]
	fmt.Println("map13===stuMap[`stu01`]", stuMap["stu01"]) //map13===stuMap[`stu01`] map[name:syf sex:girl]
}

func map14() {
	var map1 = make(map[string]string)
	map1["no1"] = "syf"
	map1["no2"] = "hello"
	fmt.Println("map14===map1:", map1) //map14===map1: map[no1:syf no2:hello]
	map1["no2"] = "smile"              //修改
	map1["no3"] = "hello"              //添加
	fmt.Println("map14===map1:", map1) //map14===map1: map[no1:syf no2:smile no3:hello]
}

func map15() {
	/*
		delete(map,"key),delete是一个内置函数，若key存在，则删除key-value,若key不存在，不操作，但也不会报错
		http://doc.golang.ltd -> builtin -> func delete(m map[Type]Type1, key Type)
	*/
	var map1 = make(map[string]string)
	map1["no1"] = "syf"
	map1["no2"] = "hello"
	delete(map1, "no2")                //key存在，则删除key-value
	delete(map1, "no3")                //key不存在，不操作，但也不会报错
	fmt.Println("map15===map1:", map1) //map15===map1: map[no1:syf]
}

func map16() {
	var map1 = make(map[string]string)
	map1["no1"] = "syf"
	map1["no2"] = "hello"
	//一次性删除所有key
	//1,遍历所有key,逐个删除
	delete(map1, "no1")
	delete(map1, "no2")
	fmt.Println("map16===map1:", map1) //map16===map1: map[]

	var map2 = make(map[string]string)
	map2["no1"] = "syf"
	map2["no2"] = "hello"
	//2,直接make一个新的空间
	map2 = make(map[string]string)
	fmt.Println("map16===map2:", map2) //map16===map2: map[]
}

func map17() {
	var map1 = make(map[string]string)
	map1["no1"] = "syf"
	map1["no2"] = "hello"
	//若map1存在key="no1",则ok=true,否则ok=false
	val, ok := map1["no1"]
	if ok {
		fmt.Println("map17===found val:", val) //map17===found val: syf
	} else {
		fmt.Println("map17===not found.")
	}
}

func map18() {
	//for-range
	var map1 = make(map[string]string)
	map1["no1"] = "syf"
	map1["no2"] = "hello"
	//map输出是无序的，map与设置值，及key,value均无关。完全无序!!!
	for k, v := range map1 {
		fmt.Printf("map18===k:%v,v:%s \n", k, v)
	}
	//双层for-range
	var stuMap = make(map[string]map[string]string)
	stuMap["stu01"] = make(map[string]string)
	stuMap["stu01"]["name"] = "syf"
	stuMap["stu01"]["sex"] = "girl"
	stuMap["stu02"] = make(map[string]string)
	stuMap["stu02"]["name"] = "smile"
	stuMap["stu02"]["sex"] = "boy"
	for k1, v1 := range stuMap {
		fmt.Printf("map18===k1:%v \n", k1)
		for k2, v2 := range v1 {
			fmt.Printf("\tmap18===k2:%v,v2:%s \n", k2, v2)
		}
	}

	fmt.Println("map18=== stuMap.len:", len(stuMap)) //map18=== stuMap.len: 2
}
