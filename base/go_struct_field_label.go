/*
 * @Author: your name
 * @Date: 2021-07-01 11:22:58
 * @LastEditTime: 2021-07-01 14:34:58
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /go_notes/base/go_struct_field_label.go
 */
package main
import (
	"encoding/json"
	"fmt"
)

type DebugInfo struct {
	Level string  `json: "level"` // level 解码为Level
	Msg  string  `json: "messgage"` // message 解码为Msg
	Author string `json: "-"`             //　忽略Author
}

func (dbgInfo DebugInfo) String() string {
	return fmt.Sprintf("{Level: %v, Msg: %v, Author: %v}", dbgInfo.Level, dbgInfo.Msg, dbgInfo.Author)
}

func main () {
	// 定义json格式字符串
	data := `[{"level": "debug", "Msg": "FileNotFound", "author": "AlexZ"},` + `{"level": "", "Msg": "Logic error", "author": "Tom"}]`
    var dbgInfos []DebugInfo
	// 将字符串按照标签解析成结构体切片
	json.Unmarshal([] byte(data), &dbgInfos)
	fmt.Println(dbgInfos)
}