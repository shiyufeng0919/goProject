package _interface

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

/*
面向对象编程思想-接口

最佳实践
1.实现对Hero结构体切片的排序: sort.Sort(data interface)
*/

func TestInterface04(t *testing.T) {
	//实现对结构体切片的排序
	interface41()
}

//声明一个Hero结构体
type Hero struct {
	Name string
	Age  int
}

//声明一个Hero结构体切片类型
type HeroSlice []Hero

//实现interface接口
//Len方法返回集合中的元素个数
func (h HeroSlice) Len() int {
	return len(h)
}

//Less方法报告索引i的元素是否比索引j的元素小。 less方法就是决定你使用什么标准进行排序。
// i < j 则升序，否则降序
func (h HeroSlice) Less(i, j int) bool {
	//标准： 按hero的年龄从小-》大 升序
	return h[i].Age < h[j].Age
}

//Swap方法交换索引i和j的两个元素
func (h HeroSlice) Swap(i, j int) {
	//完成一个交换任务,等价于下述简化版
	//temp := h[i]
	//h[i] = h[j]
	//h[j] = temp

	//简化版
	h[i], h[j] = h[j], h[i]
}

func interface41() {
	var heroSlice HeroSlice
	for i := 0; i < 6; i++ {
		hero := Hero{
			Name: fmt.Sprintf("syf%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heroSlice = append(heroSlice, hero)
	}
	fmt.Printf("interface41===\n heroSlice:%+v \n", heroSlice) //heroSlice:[{Name:syf81 Age:87} {Name:syf47 Age:59} {Name:syf81 Age:18} {Name:syf25 Age:40} {Name:syf56 Age:0} {Name:syf94 Age:11}]
	sort.Sort(heroSlice)
	fmt.Printf("sort heroSlice:%+v", heroSlice) //sort heroSlice:[{Name:syf56 Age:0} {Name:syf94 Age:11} {Name:syf81 Age:18} {Name:syf25 Age:40} {Name:syf47 Age:59} {Name:syf81 Age:87}]
}
