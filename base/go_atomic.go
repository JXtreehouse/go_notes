//Go语言在设计上对同步（Synchronization，数据同步和线程同步）提供大量的支持，比如 goroutine和channel同步原语，库层面有
//
//
//- sync：提供基本的同步原语（比如Mutex、RWMutex、Locker）和 工具类（Once、WaitGroup、Cond、Pool、Map）
//- sync/atomic：提供变量的原子操作（基于硬件指令 compare-and-swap）
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	x  int64
	mx sync.Mutex
	wg sync.WaitGroup
)

// 普通函数， 并发不安全
func Add() {
	x++
	wg.Done()
}

// 互斥锁, 并发安全，性能低于原子操作
func MxAdd() {
	mx.Lock()
	x++
	mx.Unlock()
	wg.Done()
}

// 原子操作，并发安全，性能高于互斥锁，只针对go中的一些基本数据类型使用
func AmAdd() {
	atomic.AddInt64(&x, 1)
	wg.Done()
}

func main() {
	// 原子操作atomic包
	// 加锁操作涉及到内核态的上下文切换， 比较耗时，代价高
	// 针对基本数据类型我们还可以使用原子操作来保证并发安全
	// 因为原子操作是go语言提供的方法，我们在用户态就可以完成，因此性能比加锁操作更好
	// go语言的原子操作由内置的库,sync/atomic完成

	start := time.Now()
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go Add() // 普通版Add函数不是并发安全的
		//go MxAdd() // 加锁版Add函数，是并发安全的, 但是加锁性能开销大
		//go AmAdd() // 原子操作版Add函数，是并发安全的，性能优于加锁版
	}

	end := time.Now()
	wg.Wait()
	fmt.Println(x)
	fmt.Println(end.Sub(start))

}
