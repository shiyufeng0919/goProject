package array

import (
	"fmt"
	"testing"
)

/*
多维数组-二维数组

案例1: 使用二维数组输出如下图形
	000000
	001000
	020300
	000000

语法: var 数组名 [size1][size2]类型，如: var array [2][3]int[][]，再赋值
size1: 一维数组元素个数
size2: 一维数组各个元素，所拥的元素个数

二维数组在内存的存在形式***

使用方式2: 直接初始化
1,声明: var 数组名[size1][size2]类型=[size1][size2]类型{{初值...},{初值...}}
2,赋值(有默认值，eg:int=0)

二维数组在在声明/定义时也对应有4种写法(与一维数组类似)
1,var arrayName [size1][size2]Type = [size1][size2]Type{{init value...},{init value...}}
2,var arrayName [size1][size2]Type = [...][size2]Type{{init value...},{init value...}}
3,var arrayName  = [size1][size2]Type{{init value...},{init value...}}
4,var arrayName  = [...][size2]Type{{init value...},{init value...}}

二维数组遍历
*/
func TestArray04(t *testing.T){
	//案例1
	array41()
	//二维数组在内存的存在形式
	array42()
	//二维数组直接初始化及遍历
	array43()
	//二维数组应用案例：定义一个二维数组，用于保存3个班,每个班5名同学成绩，并求出每个班平均分及所有班平均分
	array44()
}

func array41(){
	//数组有4个元素，每个元素又是一个包含6个元素的数组
	var arr [4][6]int //先声明  <=> var arr [4][6]int[][]
	fmt.Printf("\n array41===arr:%+v \n",arr)
	arr[1][2] = 1   //再赋值
	arr[2][1] = 2
	arr[2][3] = 3

	//遍历二维数组，输出要求图形
	for i:=0 ; i< 4; i++ {
		for j:=0 ; j < 6 ; j ++ {
			fmt.Print(arr[i][j])
		}
		fmt.Println()
	}
}

func array42(){
	/*
	二维数组在内存的布局
	var array [2][3]int 在内存中会有一个变量array

	array     |  指针(0xc00001c180)   |   指针(0xc00001c198)   |     -> array有2个元素，每一个元素指向了一个一维数组

	0xc00001c180 -						   0xc00001c198 -
	              |										  |
				  ^										  ^
	          |  0  |   0  |   0  |             		|  0  |   0  |   0  |
																  |
															  被修改为888

	0xc00001c198 - 0xc00001c180 = 18个字节

0xc00001c180
	0  			0   	   0    每个int=8byte
	0  			0   	   0
0xc00001c198
	*/
	var array [2][3]int
	array[1][1] = 888
	fmt.Printf("\n array42=== array:%+v",array) //[[0 0 0] [0 888 0]]
	fmt.Printf("\n array42=== array[0]的地址:%p",&array[0]) //0xc00001c180
	fmt.Printf("\n array42=== array[1]的地址:%p",&array[1]) //0xc00001c198
	fmt.Printf("\n array42=== array[0][0]的首地址:%p",&array[0][0]) //0xc00001c180
	fmt.Printf("\n array42=== array[1][0]的首地址:%p",&array[1][0]) //0xc00001c198
}

func array43(){
	var array [2][3]int = [2][3]int{{1,2,3},{4,5,6}}
	fmt.Printf("\n array43===array:%+v \n",array)
	//遍历1
	for i:=0; i<len(array); i++ {
		for j:=0 ; j< len(array[i]); j++ {
			fmt.Printf("%v\t",array[i][j])
		}
		fmt.Println()
	}

	//遍历2
	for _,v := range array {
		for _,v2 := range v {
			fmt.Printf("%v\t",v2)
		}
		fmt.Println()
	}
}

//定义一个二维数组，用于保存3个班,每个班5名同学成绩，并求出每个班平均分及所有班平均分
func array44(){
	//定义一个二维数组
	var scores [3][5]float64 = [3][5]float64{{10,20,30,40,50},{60,70,80,90,100},{10,30,50,70,90}}
	//for i:=0 ; i< len(scores); i++ {
	//	for j:=0 ; j< len(scores[i]); j++ {
	//		fmt.Printf("请输入第%d班的第%d个学生成绩\n",i+1,j+1)
	//		fmt.Scanln(&scores[i][j])
	//	}
	//}
	fmt.Printf("\n array44=== scores:%v",scores) //[[10 20 30 40 50] [60 70 80 90 100] [10 30 50 70 90]]

	//统计所有班级分数和
	var totalSum = 0.0
	var totalPerson = 0
	for i:=0 ; i< len(scores); i++ {
		//统计每个班分数和
		sum := 0.0
		for j:=0 ; j< len(scores[i]); j++ {
			sum += scores[i][j]
		}
		totalSum += sum
		totalPerson += len(scores[i])
		fmt.Printf("\n array44=== 第%d班级总分为%v,平均分:%v\n",i+1,sum,sum/float64(len(scores[i])))
	}
	fmt.Printf("\n array44=== 所有班级总分为:%v,总人数为：%d,所有班级平均分:%v",totalSum,totalPerson,fmt.Sprintf("%.2f",totalSum/float64(totalPerson)))
}
