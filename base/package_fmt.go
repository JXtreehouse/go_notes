/*
 * @Author: your name
 * @Date: 2021-05-25 16:52:28
 * @LastEditTime: 2021-05-25 17:13:24
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /go_notes/base/package_fmt_%v_%T.go
 */
package base

import (
	"fmt"
)

type Power struct {
	age int
	hight int
	name string
}

func main() {
	var i Power = Power{age:10, hight:178, name: "NewMan"}

	fmt.Printf("type:%T\n", i)
	fmt.Printf("value:%v\n", i)
	fmt.Printf("value+:%+v\n", i)
	fmt.Printf("value#:%#v\n", i)

	fmt.Println("========interface========")
	var interf interface{} = i
	fmt.Printf("%v\n", interf)
	fmt.Println(interf)
}// output:
// type:main.Power
// // output:
// type:main.Power
// value:{10 178 NewMan}
// // output:
// type:main.Power
// value:{10 178 NewMan}
// value+:{age:10 hight:178 name:NewMan}
// // output:
// type:main.Power
// value:{10 178 NewMan}
// value+:{age:10 hight:178 name:NewMan}
// value#:main.Power{age:10, hight:178, name:"NewMan"}
// // output:
// type:main.Power
// value:{10 178 NewMan}
// value+:{age:10 hight:178 name:NewMan}
// value#:main.Power{age:10, hight:178, name:"NewMan"}
// ========interface========
// {10 178 NewMan}
// {10 178 NewMan}
// 