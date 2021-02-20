package main

import (
	"bufio"
	"fmt"
	"os"
)

func printInput(ch chan string) {
	// 使用 for 循环从 channel 中读取数据
	for val := range ch {
		// 读取到结束符号
		if val == "EOF" {
			break
		}
		fmt.Printf("Input is %s\n", val)
	}
}

func main() {
	// 创建一个无缓冲的 channel
	ch := make(chan string)
	go printInput(ch)

	// 从命令行读取输入
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		val := scanner.Text()
		ch <- val
		if val == "EOF" {
			fmt.Println("End eth game!")
			break
		}
	}
	// 程序最后关闭ch
	defer close(ch)
}
