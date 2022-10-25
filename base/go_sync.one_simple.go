package main

import (
	"fmt"
	"sync"
)

var (
	once sync.Once
	n    *Woman
)

//func main() {
//
//	onceFunc := func() {
//		fmt.Println("Only once")
//	}
//
//	for i := 0; i < 10; i++ {
//		once.Do(onceFunc)
//	}
//}

// nancy的淘宝客户端和web端都会指向nancy这一个人
type Woman struct {
	name string
}

func GetNancy(i int) *Woman {

	getNameOnceFunc := func() {
		n = new(Woman)
		n.name = "Nancy"
		fmt.Println("newNancy")
	}

	fmt.Println("---start---")
	fmt.Println(i)
	fmt.Println("greetingNancy")
	once.Do(getNameOnceFunc)
	fmt.Println("---end----")
	return n
}

func main() {
	for i := 0; i < 5; i++ {
		_ = GetNancy(i)
	}
}
