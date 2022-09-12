package model

import "fmt"

/*
模拟封装
*/

//其它包不能直接访问person
type person struct {
	Name   string
	age    int //其它包不能直接访问age,salary
	salary float64
}

//写一个工厂模式的函数，相当于构造函数
func NewPerson(name string) *person {
	return &person{Name: name}
}

//为了访问age和salary编写一个SetXxx和GetXxx方法
func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("age 范围 error.")
	}
}

func (p *person) SetSalary(salary float64) {
	if salary >= 3000 && salary < 30000 {
		p.salary = salary
	} else {
		fmt.Println("salary 范围 error.")
	}
}

func (p *person) GetAge() int {
	return p.age
}

func (p *person) GetSalary() float64 {
	return p.salary
}
