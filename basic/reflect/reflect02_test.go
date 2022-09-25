package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

/*
反射注意事项和细节说明:

1,reflect.Value.Kind,获取变量的类别，返回的是一个常量。

2,Type是类型，kind是类别。Type和Kind可能是相同的，也可能是不同的。
如: var num int=99  //num的Type是int,Kind也是int
如: var stu Student //stu的Type是包名.Student，Kind是Student (Kind范围要大于Type)

3,通过反射可以让变量在interface{}和Reflect.Value之间相互转换，在代码中体现:
变量 <---> interface{} <----> reflect.Value

4,使用反射方式来获取变量的值(并返回对应的类型)，要求数据类型匹配，如: x 是 int,
则要使用reflect.Value(x).Int()，而不能使用其它的，否则Panic.

5,通过反射来修改变量，注意当使用SetXxx()方法来设置，需要通过对应的指针类型来完成，这样才能改变传入的变量的值。
同时需要使用到reflect.Value.Elem()方法.

6,reflect.Value.Elem()应该如何理解? 参见:reflect21()
*/
func TestReflect02(t *testing.T) {
	//通过反射机制，修改num int的值 (Elem()的使用!!!)
	reflect21()
}

func reflect21() {
	var num int = 888
	reflectTest211(num)
	reflectTest212(&num)
	reflectTest213(&num)
	fmt.Println("num:", num)
}

func reflectTest211(i interface{}) {
	//获取到reflect.Value
	refValue := reflect.ValueOf(i)
	fmt.Printf("refValue type=%T,kind=%v\n", refValue, refValue.Kind()) //refValue type=reflect.Value,kind=int
	//refValue.SetInt(999) //运行时错误panic: reflect: reflect.Value.SetInt using unaddressable value
	fmt.Println("----------reflectTest211 end.-----------")
}

func reflectTest212(i interface{}) {
	//获取到reflect.Value
	refValue := reflect.ValueOf(i)
	fmt.Printf("refValue type=%T,kind=%v\n", refValue, refValue.Kind()) //refValue type=reflect.Value,kind=ptr
	//refValue.SetInt(999) //运行时错误panic: reflect: reflect.Value.SetInt using unaddressable value
	fmt.Println("----------reflectTest212 end.-----------")
}

func reflectTest213(i interface{}) {
	//获取到reflect.Value
	refValue := reflect.ValueOf(i)
	fmt.Printf("refValue type=%T,kind=%v\n", refValue, refValue.Kind()) //refValue type=reflect.Value,kind=ptr
	/*
		注意：Elem()用法
		func (v Value) Elem() Value
		Elem返回v持有的接口保管的值的Value封装，或者v持有的指针指向的值的Value封装。如果v的Kind不是Interface或Ptr会panic；如果v持有的值为nil，会返回Value零值。
		refValue.Elem()类似于:
		num := 888
		ptr *int = &num
		num2 := *ptr
	*/
	refValue.Elem().SetInt(999)
	fmt.Println("----------reflectTest213 end.-----------")
}
