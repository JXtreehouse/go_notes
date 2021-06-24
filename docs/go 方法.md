<!--
 * @Author: your name
 * @Date: 2021-06-24 20:36:36
 * @LastEditTime: 2021-06-24 20:39:54
 * @LastEditors: Please set LastEditors
 * @Description: [In User Settings Edit](https://github.com/skywind3000/awesome-cheatsheets/blob/master/languages/golang.go)
 * @FilePath: /docs/go 方法.md
-->
```
//方法通常是针对一个结构体来说的

type Foo struct {
    a int
}

// 值接收者
func (f Foo) test() {
    f.a  = 1    // 不会改变原来的值
}

// 指针接收者
func (f *Foo) test() {
    f.a = 1  会改变原值
}

```
