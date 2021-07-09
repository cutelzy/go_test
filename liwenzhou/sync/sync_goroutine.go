package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Counter interface {
	Inc()
	Load() int64
}

// 普通版
type CommonCounter struct {
	counter int64
}

func (c CommonCounter) Inc() {
	c.counter++
}

func (c CommonCounter) Load() int64 {
	return c.counter
}

// 互斥锁板
type MutexCounter struct {
	counter int64
	lock    sync.Mutex
}

func (mc *MutexCounter) Inc() {
	mc.lock.Lock()
	defer mc.lock.Unlock()
	mc.counter++
}

func (mc *MutexCounter) Load() int64 {
	mc.lock.Lock()
	defer mc.lock.Unlock()
	return mc.counter
}


// 原子操作版
type AtomicCounter struct {
	couter int64
}

func (ac *AtomicCounter) Inc() {
	atomic.AddInt64(&ac.couter, 1)
}

func (ac *AtomicCounter) Load() int64 {
	return atomic.LoadInt64(&ac.couter)
}

func test(c Counter) {
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			c.Inc()
			wg.Done()	
		}()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(c.Load(), end.Sub(start))
}

func main() {
	c1 := CommonCounter{}	// 非并发安全
	test(c1)
	c2 := MutexCounter{}	// 	使用互斥锁实现并发安全
	test(&c2)
	c3 := AtomicCounter{}	// 并发安全且比互斥锁效率更高
	test(&c3)
}