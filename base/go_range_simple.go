/*
 * @Author: your name
 * @Date: 2021-08-20 14:28:32
 * @LastEditTime: 2021-08-20 14:48:58
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /go_notes/base/go_range_simple.go
 */
package main

import "fmt"

func main() {
	// 新建一个slice
	nums := []int{1,2,3,4};
	length := 0;
	for i, num := range nums {
		length++;
		fmt.Println("index", i)
		fmt.Println("value", num)
	}
	fmt.Println(length);
}
