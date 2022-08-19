/*
 * @Author: your name
 * @Date: 2021-07-01 14:53:12
 * @LastEditTime: 2021-07-01 15:05:20
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /go_notes/base/go_json_anonymous_field.go
 */
package bufferChannel

import (
	"encoding/json"
	"fmt"
)

type Point struct{ X, Y int }
type Circle struct {
	Point
	Radius int
}

func main() {
	//定义JSON格式字符串
	data := `{"X": 80, "Y": 80, "Radius": 40}`
	var c Circle
	// 将字符串解析成匿名字段
	json.Unmarshal([]byte(data), &c)
	fmt.Println(c)
}
