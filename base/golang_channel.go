/*
 * @Author: your name
 * @Date: 2021-05-26 15:56:38
 * @LastEditTime: 2021-05-26 16:13:08
 * @LastEditors: Please set LastEditors
 * @Description:
    https://halfrost.com/go_channel/
    https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/02.7.md
 * @FilePath: /go_notes/base/golang_channel.go
*/
package main

import "fmt"

func sum(s []int, c chan int) {
	fmt.Println("c1:", c)
	sum := 0
	for _, v := range s {
		sum += v
	}
	fmt.Println("sum:", sum)
	//fmt.Println(sum)
	c <- sum // 把 sum 发送到通道 c
	fmt.Println("c2:", c)
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	fmt.Println("s[:len(s)/2]:", s[:len(s)/2])
	go sum(s[:len(s)/2], c) // 启动一个 goroutine 来处理 sum()
	go sum(s[len(s)/2:], c)
	fmt.Println("c3", c)
	x, y := <-c, <-c // 从通道 c 中接收

	fmt.Println("x,y,x+y:", x, y, x+y)
}

// output:
// -5
// 17
// -5 17 12
//
