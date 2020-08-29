// 结构体实现接口
// 实现一个接口，需要满足两个条件:
// (1) 类型中添加的方法签名和接口接口签名完全一致，包括名称、参数列表、返回参数等。
// (2) 接口的所有方法均被实现。
package main
import "fmt"

// Cat 接口
type Cat interface {
	// 抓老鼠
	CatchMouse()
}

// Dog 接口
type Dog interface {
	// 犬叫
	Bark()
}

type CatDog struct {
	Name string
}

// 实现 Cat 接口
func (catDog *CatDog) CatchMouse() {
	fmt.Printf("%v caught the mouse and ate it!\n", catDog.Name)
}

// 实现Dog 接口
func (catDog *CatDog) Bark() {
	fmt.Printf("%v barked loudly!\n", catDog.Name)
}

func main() {
	catDog := &CatDog {
		"Lucy",
	}

	// 声明一个 Cat 接口，并将 catDog 指针类型赋值给cat
	var cat Cat
	cat = catDog
	cat.CatchMouse()

	// 声明一个 Dog 接口，并将 catDog dog
	var dog Dog
	dog = catDog
	dog.Bark()
}