package model

import "fmt"

/*
模拟继承
*/
type student struct {
	Name    string
	age     int
	score   float64
	Content string
}

type Class struct {
	Content string
}

//小学生
type Pupil struct {
	student        //嵌入了student匿名结构体，即该结构体没有给名字.
	Name    string //当结构体和匿名结构体有相同的字段或方法时，编译器采用就近访问原则访问。若希望访问匿名结构体的字段和方法，可通过匿名结构体名来区分。
	Class          //结构体嵌入2/多个匿名结构体，若两个匿名结构体有相同的字段和方法(同时结构体本身没有同名的字段和方法)，在访问时，就必须明确指定匿名结构体名字，否则编译器报错
}

//大学生
type Graduate struct {
	student
}

func (stu *student) ShowInfo() {
	fmt.Println("stu:", *stu)
}

func (stu *student) SetCores(score float64) {
	if score > 0 && score <= 100 {
		stu.score = score
	} else {
		fmt.Println("set score is error.")
	}
}

func (stu *student) SetAge(age int) {
	if age > 0 && age < 150 {
		stu.age = age
	} else {
		fmt.Println("set age is error.")
	}
}

func (stu *student) getScore() float64 {
	return stu.score
}

func (stu *student) getAge() int {
	return stu.age
}

func (p *Pupil) Test() {
	fmt.Println("小学生考试测验...")
}

func (p *Graduate) Test() {
	fmt.Println("大学生考试测验...")
}

func (p *Pupil) ReturnScore() {
	fmt.Println("return score:", p.getScore())
}
