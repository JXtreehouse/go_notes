package main

import "fmt"

func observe(i interface{}) {
	// using the format specifier
	// %T to check type in interface
	fmt.Printf("The type passed is: %T\n", i)

	// using the format Specifier %#v
	// to check value in interface
	fmt.Printf("The value passed is: %#v \n", i)

	fmt.Println("===============================")
}

func main() {
	var value float64 = 15
	value2 := "youzi programmer"
	observe(value)
	observe(value2)
}

//在这里我们可以清楚地看到一个空接口将能够接受任何参数并适应它的值和数据类型。这包括但不限于结构和指向结构的指针。
