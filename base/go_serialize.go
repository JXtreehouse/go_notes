/*
 * @Author: your name
 * @Date: 2021-06-29 16:59:34
 * @LastEditTime: 2021-06-29 17:47:43
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /go_notes/base/go_serialize.go
 */
package base

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name string `json: "monster_name"`
	Age int
   Birthday string
   Sal string
   Skill string
}

func main() {
	monster := Monster {
		Name: "哥斯拉",
		Birthday: "1999-11-11",
		Age: 20,
		Sal: "1000",
		Skill: "coding & loving"
	}

	data, err := json.Marshal(&monster)
	// 序列化
	// json.Unmarshal([]byte(str), &monster)
	 if err != nil {
		 fmt.Print("序列化错误 err=%v", err)
	 }

	 fmt.Printf("序列化=%v\n", string(data))

}

func testMap() {
	// key 是string value 为任意类型
	var a map [string] interface{}
	a = make(map [string] interface{})
	a["name"] = "哥斯拉"
	a["age"] = 20

	// map 本身就是引用类型
	data, err := json.Marshal(a)
	// 反序列化
	// 注意：反序列化 map,不需要 make,因为 make 操作被封装到 Unmarshal 函数 
	//err := json.Unmarshal([]byte(str), &a)
	if err!=nil {
		fmt.Print("序列化错误err=%v", err)
	}
	fmt.Printf("序列化后=%v\n", string(data))　
}

func testSlice()  {
	var slice []map [string] interface{}
	var m1 map [string] interface{}

	m1 = make(map[string] interface{})
	m1["name"] = "jack"
	m1["age"] = 20
	slice = append(slice, m1)

	var m2 map[string]interface{}

	m2 = make(map[string]interface{})
	m2["name"] = "tom"
	m2["age"] = 21
	slice = append(slice, m2)

	data, err := json.Marshal(slice)

	//反序列化，不需要 make,因为 make 操作被封装到 Unmarshal 函数 
	//err := json.Unmarshal([]byte(str), &slice)

	if err != nil {
		fmt.Printf("序列化错误 err=%v", err)
	}

	fmt.Printf("序列化后=%v\n", string(data))
}