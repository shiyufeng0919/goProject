package _interface

import "testing"

/*
接口 vs 继承

实现接口可以看作是对继承的一种补充。

					    learning english skill
							interface
        运动员				 /		\			   Student
      /      \				/	 	 \			/			\
basketball   football   <--   			--> 	big Student   small Student

需求: 要求运动员和Student都能撑握学英语技能.

不能在运动员或student处增加english，因为运动员所要撑握的英语技能与学生可能是不样的，各有各的规范。

一。接口和继承解决的问题不同：
1，继承的价值主要在于：解决代码复用性和可维护性
2，接口的价值主要在于：「设计」，设计好各种「规范(方法)」，让其它自定义类型去实现这些方法

二。接口比继承更加灵活
继承是满足 is -a 的关系，而接口只需满足like - a的关系

is -a关系： Person{}  Student{} -> student is a person.

like - a关系：Bird{}  Monkey{}  -> monkey like a bird fly.

三。接口在一定程序上实现代码解耦
*/
func TestInterface06(t *testing.T) {

}
