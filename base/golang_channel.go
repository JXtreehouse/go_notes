/*
 * @Author: your name
 * @Date: 2021-05-26 15:56:38
 * @LastEditTime: 2021-05-26 16:13:08
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /go_notes/base/golang_channel.go
 */
 package base

 import "fmt"
 
 func sum(s []int, c chan int) {
		 sum := 0
		 for _, v := range s {
				 sum += v
		 }
		 fmt.Println(sum)
		 c <- sum // 把 sum 发送到通道 c
 }
 
 func main() {
		 s := []int{7, 2, 8, -9, 4, 0}
 
		 c := make(chan int)
		//  fmt.Println(s[:len(s)/2])
		 go sum(s[:len(s)/2], c)
		 go sum(s[len(s)/2:], c)
		 x, y := <-c, <-c // 从通道 c 中接收
 
		 fmt.Println(x, y, x+y)
 }

// output:
// -5
// 17
// -5 17 12
// 