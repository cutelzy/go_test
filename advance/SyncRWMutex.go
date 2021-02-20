package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// 读写锁
var rwLock sync.RWMutex

func main() {

	// 获取读锁
	for i := 0; i < 5; i++ {
		go func(i int) {
			rwLock.RLock()
			defer rwLock.RUnlock()
			fmt.Println("read func " + strconv.Itoa(i) + " get rlock at " + time.Now().String())
			time.Sleep(time.Second)
		}(i)
	}

	time.Sleep(time.Second / 10)
	// 获取写锁
	for i := 0; i < 5; i++ {
		go func(i int) {
			rwLock.Lock()
			defer rwLock.Unlock()
			fmt.Println("write func " + strconv.Itoa(i) + " get wlock at " + time.Now().String())
			time.Sleep(time.Second)
		}(i)
	}

	// 保证所有的 goroutine 执行结束
	time.Sleep(time.Second * 10)
}
