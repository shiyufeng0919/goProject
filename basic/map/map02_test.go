package _map

import (
	"fmt"
	"sort"
	"testing"
)

/*
一。map切片
var slice []map
切片的数据类型若是map,则称为slice of map. map切片这样使用则map个数就可以「动态变化」了.

二。map排序
1,golang中没有一个专门方法针对map的key进行排序
2,golang中map默认是无序的，注意不是按照添加的顺序存放，每次遍历map,输出可能不一样。
3,golang中的map排序，是先将key进行排序，然后中根据key值遍历输出即可.

三。map使用细节
1,map是引用类型，遵守引用类型传递机制，在一个函数接收map,修改后，会直接修改原来的map
2,map的容量达到后，再想在map增加元素，会自动扩容，并不会panic，即map能动态的增长key-value
3,map的value也经常使用struct类型，更适合管理复杂的数据(eg:比value是一个map更好)
*/

func TestMap02(t *testing.T) {
	//使用一个map来记录monster的信息name和age,也就是说一个monster对应一个map,并且monster的个数可以动态的增加 => map切片
	//演示map切片
	map21()
	//map排序(按key排序),借助slice
	map22()
	//map是引用类型，遵守引用类型传递机制，在一个函数接收map,修改后，会直接修改原来的map
	map23()
	//map的容量达到后，再想在map增加元素，会自动扩容，并不会panic，即map能动态的增长key-value (切片不可以，需要使用append)
	map24()
	//map的value也经常使用struct类型，更适合管理复杂的数据(eg:比value是一个map更好)
	map25()
}

func map21() {
	//事先分配了2个空间，只能放置2个map[string]string
	monsters := make([]map[string]string, 2)
	//未分配数据空间
	if monsters[0] == nil {
		//分配数据空间
		monsters[0] = make(map[string]string)
		monsters[0]["name"] = "syf"
		monsters[0]["sex"] = "girl"
	}
	if monsters[1] == nil {
		//分配数据空间
		monsters[1] = make(map[string]string)
		monsters[1]["name"] = "smile"
		monsters[1]["sex"] = "girl"
	}
	//panic: runtime error: index out of range [2] with length 2 [recovered]
	//if monsters[2] == nil {
	//	//分配数据空间
	//	monsters[2] = make(map[string]string)
	//	monsters[2]["name"] = "hello"
	//	monsters[2]["sex"] = "body"
	//}

	//解决方案: 通过map切片append函数动态增加slice
	monster := map[string]string{
		"name": "hello",
		"sex":  "boy",
	}
	monsters = append(monsters, monster)
	fmt.Printf("map21===monsters:%+v", monsters) //map21===monsters:[map[name:syf sex:girl] map[name:smile sex:girl] map[name:hello sex:boy]]
}

func map22() {
	var map1 map[int]int
	map1 = make(map[int]int, 10)
	map1[99] = 199
	map1[88] = 188
	map1[100] = 100
	map1[66] = 166

	/*
		若按照map的key的顺序进行排序输出，则
		1,先将map的key放入到切片中
		2,对切片排序
		3,遍历切片，按照key输出map值
	*/
	var keys []int
	for k := range map1 { //<=> for k,_ := range map1
		keys = append(keys, k)
	}
	fmt.Println("\nmap22===keys:", keys) //map22===keys: [88 66 99 100]
	//排序
	sort.Ints(keys)
	fmt.Println("map22===order keys:", keys) //map22===order keys: [66 88 99 100]

	for _, v := range keys {
		fmt.Printf("map22=== map[%v]=%v \n", v, map1[v])
	}
}

func map23() {
	map1 := make(map[int]int)
	map1[0] = 100
	map1[1] = 99
	//map是引用类型
	modifyMap(map1)
	fmt.Printf("map23=== map1:%+v \n", map1) //map23=== map1:map[0:999 1:99]
}

func map24() {
	map1 := make(map[int]int, 1)
	map1[0] = 100
	map1[1] = 99                             //给了2个值也不会panic,会动态扩容
	fmt.Printf("map24=== map1:%+v \n", map1) //map24=== map1:map[0:100 1:99]
}

type stu struct {
	Name string
	Sex  string
}
func map25() {
	/*
		改造map13()示例,student使用map存储
		map的key是学号
		map的value是学生结构体
	*/
	students := make(map[string]stu)
	students["no1"] = stu{
		Name: "syf",
		Sex:  "girl",
	}
	students["no2"] = stu{
		Name: "smile",
		Sex:  "body",
	}
	fmt.Printf("map25===students:%+v \n", students)

	for k,v := range students {
		fmt.Printf("map25=== students no:%s \n",k)
		fmt.Printf("map25=== students name:%s \n",v.Name)
		fmt.Printf("map25=== students sex:%s \n",v.Sex)
	}
}

func modifyMap(map1 map[int]int) {
	//对其key=0的value做修改，会影响函数外部
	map1[0] = 999
}
