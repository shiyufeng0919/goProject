package _interface

import (
	"fmt"
	"testing"
)

/*
接口 vs  继承

old monkey             bird(会飞)             fish(游)
    |
	| extends			| implement			  | implement
	^
small monkey  ----------  --------------------

通过实现方式增加small monkey的功能。即接口是对继承的一个补充。

小结:
1,当一个A结构体继承了另外一个B结构体，那么A结构体就自动继承了B结构体的字段的方法，并且可以直接使用。
2,当A结构体需要扩展功能，同时又不希望破坏继承关系，则可以去实现某个接口即可。因此我们可以认为，实现接口是对继承机制的补充。
*/

func TestInterface05(t *testing.T) {
	//接口和继承关系
	interface51()
}

//定义一个Monkey结构体
type Monkey struct {
	Name string
}

//定义Bird接口，方法会fly
type Bird interface {
	Flying()
}

//定义Fish接口，方法会swimming
type Fish interface {
	Swimming()
}

//定义Monkey的行为会爬树
func (this *Monkey) climbing() {
	fmt.Println(this.Name, "生来会爬树")
}

type LittleMonkey struct {
	Monkey //继承Monkey，则拥有Monkey下的所有属性和方法
	//该结构体想要扩展一下功能.eg:像bird一样fly,像fish一样swimming. (注意：只是该结构体想扩展，并不是所有Monkey都要扩展!!!)
	//但monkey本来就不会fly,不会swimming,因此不能做为Monkey的方法。即不能修改原Monkey.
	//借助接口实现扩展功能!!!
}

//仅针对LittleMonkey这个结构体做扩展，不影响原Monkey
func (this *LittleMonkey) Flying() {
	fmt.Println(this.Name + "通过学习会fly.")
}

func (this *LittleMonkey) Swimming() {
	fmt.Println(this.Name + "通过学习会swimming.")
}

func interface51() {
	//创建一个LittleMonkey实例
	var littleMonkey LittleMonkey
	littleMonkey.Name = "孙悟空"
	littleMonkey.climbing() //孙悟空 生来会爬树

	//实现接口在不破坏原先继承的方式下，对我们的结构体进行功能扩展。(解耦方式)
	littleMonkey.Flying()
	littleMonkey.Swimming()
}
