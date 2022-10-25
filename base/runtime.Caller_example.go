/*
 * @Author: AlexZ33
 * @Date: 2022-04-26 14:55:01
 * @LastEditTime: 2022-04-26 15:09:52
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_notes/base/runtime.Caller_example.go
 */

package main

import (
	"fmt"
	"runtime"
)

func main() {
	for i := 0; i < 4; i++ {
		test(i)
	}
}

func test(skip int) {
	call(skip)
}

func call(skip int) {
	pc, file, line, ok := runtime.Caller(skip)
	pcName := runtime.FuncForPC(pc).Name() //获取函数名
	fmt.Println(fmt.Sprintf("%v   %s   %d   %t   %s", pc, file, line, ok, pcName))
}
