//Person.go    // 方法和接收器（指针类型接收器、非指针类型接收器）
package main
import "fmt"

type Person struct {
	Name string // 姓名
	Birth string  // 生日
	ID	int64  // 身份证号
}

// 指针类型，修改个人信息
func (person *Person) changeName(name string) {
	person.Name = name
}

// 非指针类型，打印个人信息
func (person Person) printMess() {
	fmt.Printf("My name is %v, and my birthday is %v, and my id is %v\n", person.Name, person.Birth, person.ID)

	// 尝试修改个人信息，但是对原接收器并没有影响
	// person.ID = 3

}

func main() {
	p1 := Person{
		Name: "王小二",
		Birth: "1992-12-23",
		ID: 1,
	}

	p1.printMess()
	p1.changeName("王老二")
	p1.printMess()

}