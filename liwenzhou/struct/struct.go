package main

import (
	"fmt"
)

type strudent struct {
	name string
	age  int
}

func main() {
	m := make(map[string]*strudent)
	stus := []strudent{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	for _, stu := range stus {
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}


// 指针引用，都指向最后一个变量的内存地址
/*
小王子 => 大王八
娜扎 => 大王八
大王八 => 大王八
*/