<!--
 * @Author: your name
 * @Date: 2021-08-20 14:27:24
 * @LastEditTime: 2021-08-20 15:05:58
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /go_notes/docs/go range(语言范围).md
-->

Go 语言中 range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对。

# 简单循环slice

```go
//go_range_simple.go
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

```

# 循环键值对

```go
// go_range_k_v.go
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
```


# 通过 range 获取参数列表: