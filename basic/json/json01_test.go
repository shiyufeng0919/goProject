package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

/*
「encoding/json」
json基本介绍
1，json(javascript object notation):是一种轻量级的数据交换格式,易于阅读和编写，及机器解析和生成.
2，json易于机器解析和生成，有效提升网络传输效率。通常程序在网络传输时会先将数据(struct/map等)序列化成json字符串，
接收方收到Json字符串时，在反序列化恢复成原来的数据类型(struct/map等)。
struct/map ->  json        (序列化)
json       ->  struct/map  (反序列化)

	   序列化		  网络传输         反序列化
golang  ->  json字符串    ->     程序    ->     其它语言

对基本数据类型序列化意义不大!!!
*/

func TestJson01(t *testing.T) {
	//struct序列化
	json11()
	//map序列化
	json12()
	//切片序列化
	json13()
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func json11() {
	var person Person
	person.Name = "syf"
	person.Age = 18
	content, err := json.Marshal(&person)
	if err != nil {
		fmt.Println("json11===>\nerr:", err)
	}
	fmt.Println("struct序列化后content:", string(content)) //struct序列化后content: {"Name":"syf","Age":18}
}

func json12() {
	var map1 = make(map[string]interface{})
	map1["name"] = "syf"
	map1["age"] = 18
	content, err := json.Marshal(&map1)
	if err != nil {
		fmt.Println("json12===>\nerr:", err)
	}
	//序列化后结果与设置map顺序无关。
	fmt.Println("map序列化后content:", string(content))
}

func json13() {
	var slice1 []map[string]interface{}
	var map1 = make(map[string]interface{})
	map1["name"] = "syf"
	map1["age"] = 18
	var map2 = make(map[string]interface{})
	map2["name"] = "双双"
	map2["age"] = 16
	slice1 = append(slice1, map1, map2)
	content, err := json.Marshal(&slice1)
	if err != nil {
		fmt.Println("json13===>\nerr:", err)
	}
	fmt.Println("slice序列化后content:", string(content))
}
