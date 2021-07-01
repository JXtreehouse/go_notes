/*
 * @Author: your name
 * @Date: 2021-06-30 14:55:47
 * @LastEditTime: 2021-07-01 11:54:20
 * @LastEditors: Please set LastEditors
 * @Description: JSON转结构体
 * @FilePath: /go_notes/base/go_json_to_struct.go
 */

 package main

 import (
	 "encoding/json"
	 "fmt"
 )

 type DebugInfo struct {
	 Level string
	 Msg string
	 Author string  　// 未导出字段不会被json解析
 }

 func (dbgInfo DebugInfo) String() string {
	 return fmt.Sprintf("{Level:  , Msg: %s}", dbgInfo.Level, dbgInfo.Msg)
 }

 func main() {
	 // 定义JSON格式字符串
	 data := `[{"level": "debug", "msg": "File not found", "Author": "alex"}]`

	 var dbgInfos []DebugInfo
	 //将字符串解析成结构体切片
	 json.Unmarshal([]byte(data), &dbgInfos)
	 fmt.Println(dbgInfos)
 }

