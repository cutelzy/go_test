package main

import (
	"fmt"
)

/*
// Person 结构体
type Person struct {
	name string
	age  int
}

// NewPerson 构造函数
func NewPerson(name string, age int) *Person {
	return &Person{
		name,
		age,
	}
}

// Dream Person 做梦方法
func (p *Person) Dream() {
	fmt.Printf("%s的梦想是学号go语言！\n", p.name)
}

func main() {
	p1 := NewPerson("jack", 18)
	p1.Dream()
}
*/


/* -----------------------------------------  */

//Int类型添加方法
//MyInt 将int定义为自定义MyInt类型
type MyInt int

//SayHello 为MyInt添加一个SayHello的方法
func (m MyInt) SayHello() {
	fmt.Println("Hello, 我是一个int。")
}

func main() {
	var m1 MyInt
	m1.SayHello()
	m1 = 100
	fmt.Printf("%#v  %T\n", m1, m1) //100  main.MyInt
}
