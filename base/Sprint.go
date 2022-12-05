package main

import "strconv"

// 写一个函数，它可以输出任意类型的任意值，甚至是用户自定义的一个类型
// 该函数只接受一个参数, 并且与fmt.Sprint一样返回一个字符串

func Sprint(x interface{}) string {
	type stringer interface {
		string() string
	}
	switch x := x.(type) {
	case stringer:
		return x.string()
	case string:
		return x
	case int:
		return strconv.Itoa(x)
	// 对int16, uint32等类型做类型处理
	case bool:
		if x {
			return "true"
		}
		return "false"
	default:
		// array、chan、func 、map 、pointer、slice、struct
		return "???"
	}

}
