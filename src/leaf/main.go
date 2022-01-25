/*
 * @Author: your name
 * @Date: 2021-11-09 20:24:42
 * @LastEditTime: 2021-11-09 20:30:38
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_notes/leaf/main.go
 */
package main

import (
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	h := &LeafSvr{}
	h.int()

	h.Run()
}