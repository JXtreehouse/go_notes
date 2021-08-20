/*
 * @Author: your name
 * @Date: 2021-08-20 14:51:55
 * @LastEditTime: 2021-08-20 15:04:05
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /go_notes/base/go_range_k_v.go
 */
package main

import "fmt"

func main () {
	nums := []int{1,2,3,4}
	for i,num := range nums {
	   fmt.Printf("索引是%d,长度是%d\n",i, num)
	}
}