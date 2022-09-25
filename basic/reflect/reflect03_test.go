package reflect

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

/*
反射的最佳实践:

1,使用反射来遍历结构体的字段，调用结构体的方法，并获取结构体标签的值
注意以下两个重要方法:
1),func (v Value) Method(i int) Value
返回v持有值类型的第i个方法的已绑定（到v的持有值的）状态的函数形式的Value封装。返回值调用Call方法时不应包含接收者；
返回值持有的函数总是使用v的持有者作为接收者（即第一个参数）。如果i出界，或者v的持有值是接口类型的零值（nil），会panic。

默认按方法名排序，对应i值，i从0开始.

2),func (v Value) Call(in []Value) []Value
Call方法使用输入的参数in调用v持有的函数。例如，如果len(in) == 3，v.Call(in)代表调用v(in[0], in[1], in[2])（其中Value值表示其持有值）。
如果v的Kind不是Func会panic。它返回函数所有输出结果的Value封装的切片。和go代码一样，每一个输入实参的持有值都必须可以直接赋值给函数对应输入参数的类型。
如果v持有值是可变参数函数，Call方法会自行创建一个代表可变参数的切片，将对应可变参数的值都拷贝到里面。
*/
func TestReflect03(t *testing.T) {
	//使用反射来遍历结构体的字段，调用结构体的方法，并获取结构体标签的值
	reflect31()
	//使用反射方式来获取结构体的tag标签，遍历字段的值，修改字段值，调用结构体方法(要求:通过传递地址的方式完成)
	reflect32()
	//定义两个函数，再定义一个适配器函数，用作统一的处理接口
	reflect33()
	//使用反射操作任意结构体类型
	reflect34()
	//使用反射创建并操作结构体
	reflect35()
}

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score float32
	Sex   string
}

//第1个方法
func (m Monster) Print() {
	fmt.Println("---start---")
	fmt.Println(m)
	fmt.Println("---end---")
}

//第2个方法
func (m Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

//第3个方法
func (m Monster) Set(name string, age int, score float32, sex string) {
	m.Name = name
	m.Age = age
	m.Score = score
	m.Sex = sex
}

func reflect31() {
	var m Monster = Monster{
		Name:  "孙悟空",
		Age:   10000,
		Score: 100,
		Sex:   "男",
	}
	reflectTest31(m)
	fmt.Println("---------reflect31 END---------------")

}

func reflectTest31(i interface{}) {
	//获取reflect.Type类型
	refType := reflect.TypeOf(i)
	//获取reflect.Value类型
	refValue := reflect.ValueOf(i)
	//获取到i对应的类别
	kind := refValue.Kind()
	//验证传入的是否为结构体
	if kind != reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	fmt.Println("获取结构体字段及Tag标签")

	//获取到该结构体有几个字段
	num := refValue.NumField()
	fmt.Printf("struct has %d fields\n", num) //struct has 4 fields

	//遍历结构体的所有字段 i:代表结构体的第几个字段;refValue.Field(i):返回结构体的第i个字段
	for i := 0; i < num; i++ {
		fmt.Printf("Field %d: 值为%v\n", i, refValue.Field(i))
		//获取到struct标签，注意需要通过reflect.Type来获取tag标签的值
		tagVal := refType.Field(i).Tag.Get("json")
		//如果该字段有tag标签则显示
		if tagVal != "" {
			fmt.Printf("Field %d: tag为%v\n", i, tagVal)
		}
	}

	fmt.Println("获取结构体方法")

	//获取到该结构体变量有多少个方法
	numOfMethod := refValue.NumMethod()
	fmt.Printf("struct has %d method\n", numOfMethod) //struct has 3 method

	fmt.Println("调用结构体方法")

	/*
		refValue.Method(1):获取到第2个方法;Call():调用(未设参数)
		此处调用第2个方法应该为GetSum(n1, n2 int)，但此处实际调用了Print()方法!!!
		注意:方法的排序默认是按照函数名排序(ASCII码大小排序)，与方法输写顺序无关。
	*/
	refValue.Method(1).Call(nil)

	//调用第一个方法(GetSum(n1, n2 int))
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := refValue.Method(0).Call(params)
	fmt.Println("res:", res[0].Int()) //res: 50
}

func reflect32() {
	var m Monster = Monster{
		Name:  "孙悟空",
		Age:   10000,
		Score: 100,
		Sex:   "男",
	}
	//Marshal就是通过反射获取到struct的tag值
	result, _ := json.Marshal(m)
	fmt.Println("json result:", string(result))
	reflectTest32(&m)
	fmt.Println("m:", m)
	fmt.Println("---------reflect32 END---------------")

}

func reflectTest32(i interface{}) {
	//获取reflect.Type类型
	refType := reflect.TypeOf(i)
	//获取reflect.Value类型
	refValue := reflect.ValueOf(i)
	//获取到i对应的类别
	kind := refValue.Kind()
	//验证传入的是否为结构体
	if kind != reflect.Ptr && refValue.Elem().Kind() == reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	fmt.Println("获取结构体字段及Tag标签")

	//获取到该结构体有几个字段
	num := refValue.Elem().NumField()
	fmt.Printf("struct has %d fields\n", num) //struct has 4 fields

	refValue.Elem().Field(0).SetString("皮皮双")

	//遍历结构体的所有字段 i:代表结构体的第几个字段;refValue.Field(i):返回结构体的第i个字段
	for i := 0; i < num; i++ {
		fmt.Printf("Field %d: 值为%v\n", i, refValue.Elem().Field(i).Kind())
	}
	tag := refType.Elem().Field(0).Tag.Get("json")
	fmt.Printf("tag=%s\n", tag)

	fmt.Println("获取结构体方法")

	//获取到该结构体变量有多少个方法
	numOfMethod := refValue.Elem().NumMethod()
	fmt.Printf("struct has %d method\n", numOfMethod) //struct has 3 method

	fmt.Println("调用结构体方法")
	refValue.Elem().Method(1).Call(nil)
}

func reflect33() {
	call1 := func(v1 int, v2 int) {
		fmt.Printf("v1:%v,v2:%v \n", v1, v2)
	}
	call2 := func(v1 int, v2 int, v3 string) {
		fmt.Printf("v1:%v,v2:%v,v3:%s \n", v1, v2, v3)
	}
	var (
		function reflect.Value
		intValue []reflect.Value
		n        int
	)
	bridge := func(call interface{}, args ...interface{}) {
		n = len(args)
		intValue = make([]reflect.Value, n)
		for i := 0; i < n; i++ {
			intValue[i] = reflect.ValueOf(args[i])
		}
		function = reflect.ValueOf(call)
		function.Call(intValue)
	}
	bridge(call1, 1, 2)
	bridge(call2, 8, 9, "syf")
	fmt.Println("---------reflect33 END---------------")
}

func reflect34() {
	var (
		model *Monster
		sv    reflect.Value
	)
	model = &Monster{}
	sv = reflect.ValueOf(model)
	fmt.Printf("reflect.ValueOf:%s\n", sv.Kind().String())
	sv = sv.Elem()
	fmt.Printf("reflect.ValueOf.Elem:%s\n", sv.Kind().String())
	sv.FieldByName("Name").SetString("syf")
	sv.FieldByName("Age").SetInt(18)
	fmt.Printf("model:%+v\n", model)
	fmt.Println("---------reflect34 END---------------")
}

func reflect35() {
	var (
		model *Monster
		st    reflect.Type
		elem  reflect.Value
	)
	//获取类型 *Monster
	st = reflect.TypeOf(model)
	fmt.Printf("reflect.TypeOf:%s\n", st.Kind().String()) //reflect.TypeOf:ptr
	//st指向的类型
	st = st.Elem()
	fmt.Printf("reflect.TypeOf.Elem:%s\n", st.Kind().String()) //reflect.TypeOf.Elem:struct
	//new返回一个value类型值，该值持有一个指向类型为type的新申请的零值的指针
	elem = reflect.New(st)
	fmt.Printf("reflect.New:%s\n", elem.Kind().String())             //reflect.New:ptr
	fmt.Printf("reflect.New.Elem:%s\n", elem.Elem().Kind().String()) //reflect.New.Elem:struct
	//model就是创建的Monster结构体变量(实例)
	//model是*user它的指向和elem是一样的
	model = elem.Interface().(*Monster)
	//获得elem指向的值
	elem = elem.Elem()
	elem.FieldByName("Name").SetString("smile")
	elem.FieldByName("Age").SetInt(19)
	fmt.Printf("model.Name:%s,model.Age:%v\n", model.Name, model.Age) //model.Name:smile,model.Age:19
	fmt.Println("---------reflect35 END---------------")
}
