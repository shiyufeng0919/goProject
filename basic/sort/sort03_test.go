package sort

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func TestSort03(t *testing.T) {
	//使用系统默认sort包对"int切片"进行排序
	sort31()
	//使用系统提供的方法对"结构体切片"进行排序 (利用: func Sort(data Interface))
	sort32()
}

func sort31() {
	var intSlice = []int{0, -1, 10, 7, 90}
	sort.Ints(intSlice)   //递增排序
	fmt.Println(intSlice) //[-1 0 7 10 90]
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
	//return h[i].Age < h[j].Age
	//标准：按hero的名称从小-》大排序
	return h[i].Name < h[j].Name
}

//Swap方法交换索引i和j的两个元素
func (h HeroSlice) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func sort32() {
	var heroSlice HeroSlice
	for i := 0; i < 3; i++ {
		hero := Hero{
			Name: fmt.Sprintf("syf%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heroSlice = append(heroSlice, hero)
	}
	fmt.Printf("sort32===\n heroSlice:%+v \n", heroSlice) // heroSlice:[{Name:syf81 Age:87} {Name:syf47 Age:59} {Name:syf81 Age:18}]
	sort.Sort(heroSlice)
	fmt.Printf("sort heroSlice:%+v", heroSlice) //sort heroSlice:[{Name:syf47 Age:59} {Name:syf81 Age:87} {Name:syf81 Age:18}]--- PASS: TestSort03 (0.00s)
}
