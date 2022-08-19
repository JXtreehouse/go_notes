/*
 * @Author: your name
 * @Date: 2021-06-29 16:33:49
 * @LastEditTime: 2021-06-29 16:42:51
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /go_notes/base/go_func.go
 */

 package bufferChannel

 import (
	 "fmt"
 )

 type A struct {
	 Num int
 }


type Person struct {
	Name string
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

//为了提高效率，通常我们方法和指针类型绑定

func (c *Circle) area2 float{
	c.radius = 10
	return 3.14 * c.radius * c.radius
}