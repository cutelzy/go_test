package main

import (
	"fmt"
)

// func main() {
// 	a := 10
// 	b := &a
// 	fmt.Printf("a:%d ptr:%p\n", a, &a) // a:10 ptr:0xc00001a078
// 	fmt.Printf("b:%p type:%T\n", b, b) // b:0xc00001a078 type:*int
// 	fmt.Println(&b)					   // 0xc00000e018
// 	fmt.Println(*b)					   // 10
// }

// func main() {
// 	//指针取值
// 	a := 10
// 	b := &a // 取变量a的地址，将指针保存到b中
// 	fmt.Printf("type of b:%T\n", b)
// 	c := *b // 指针取值（根据指针去内存取值）
// 	fmt.Printf("type of c:%T\n", c)
// 	fmt.Printf("value of c:%v\n", c)
// }


/* ------------------------------ */
func modify1(x int) {
	x = 100
}

func modify2(x *int) {
	*x = 100
}

func main() {
	a := 10
	modify1(a)
	fmt.Printf("modify1: %d\n", a)	// 10
	modify2(&a)
	fmt.Printf("modify2: %d\n", a)	// 100
}