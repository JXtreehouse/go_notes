/*
 * @Author: your name
 * @Date: 2021-08-20 15:05:26
 * @LastEditTime: 2021-08-20 15:11:38
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /go_notes/base/go_range_param.go
 */
 package main

 import (
	 "fmt"
	 "os"
 )

 func main() {
	 fmt.Println(len(os.Args))
	 for _, arg := range os.Args{
		 fmt.Println(arg)
	 }
 }