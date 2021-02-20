package main

import (
	"fmt"
	"time"
)

func consume(ch chan int) {
	// 线程休息 100s 再从 channel 读取数据
	time.Sleep(time.Second * 100)
	<-ch

}

func main() {
	// 创建一个长度为 2 的channel
	ch := make(chan int, 2)
	go consume(ch)

	ch <- 0
	ch <- 1
	// 发送数据不被阻塞
	fmt.Println(("I am free!"))
	ch <- 2
	fmt.Println("I can not go there within 100s!")

	time.Sleep(time.Second)
}
