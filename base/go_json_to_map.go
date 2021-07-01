/*
 * @Author: your name
 * @Date: 2021-06-30 14:28:13
 * @LastEditTime: 2021-06-30 14:43:22
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /go_notes/base/go_json_to_map.go
 */

 package main
 
 import (
	 "encoding/json"
	 "fmt"
 )

 func main() {
	 // 定义JSON格式的字符串
	 data := `[{"Level": "debug", "Msg": "xxxxxx"}]`
	 var dbgInfos []map[string] string
	 // 将字符串解析成map切片
	 json.Unmarshal([] byte(data), &dbgInfos)
	 fmt.Println(dbgInfos)
 }
