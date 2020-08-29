//Pointer.go
package main
import "fmt"

func main() {
	// 声明一个 string 类型
	str := "Golang is Good!"

	// 获取 str 的指针
	strPrt := &str

	// 获取 strPrt 的指针
	// strPrtPrt := &strPrt
	// fmt.Printf("strPrtPrt type is %T, and value is %v\n", strPrtPrt, strPrtPrt)

	fmt.Printf("str type is %T, and value is %v\n", str, str)
	fmt.Printf("strPrt type is %T, and value is %v\n", strPrt, strPrt)

	newStr := *strPrt	//获取指针对应变量的值
	fmt.Printf("newStr type is %T, value is %v, and address is %p\n", newStr, newStr, &newStr)

	*strPrt = "Java is Good too!"
	fmt.Printf("newStr type is %T, value is %v, and address is %p\n", newStr, newStr, &newStr)
	fmt.Printf("str type is %T, value is %v, and address is %p\n", str, str, &str)
}