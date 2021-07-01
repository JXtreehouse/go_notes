<!--
 * @Author: your name
 * @Date: 2021-06-24 20:36:36
 * @LastEditTime: 2021-06-29 16:53:10
 * @LastEditors: Please set LastEditors
 * @Description: [In User Settings Edit](https://github.com/skywind3000/awesome-cheatsheets/blob/master/languages/golang.go)
 * @FilePath: /docs/go 方法.md
-->
# 方法的定义

```
func (recevier type) methodName(参数列表) (返回值列表){
    方法体
  return 返回值
}
```
- 1.参数列表:表示方法输入
- 2.recevier type : 表示这个方法和 type 这个类型进行绑定，或者说该方法作用于 type 类型
- 3.receiver type : type 可以是结构体，也可以其它的自定义类型
- 4.receiver : 就是 type 类型的一个变量(实例)，比如 :Person 结构体 的一个变量(实例)
- 5.返回值列表:表示返回的值，可以多个
- 6.方法主体:表示为了实现某一功能代码块
- 7.return 语句不是必须的。
# 方法通常是针对一个结构体来说的
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

# 方法和指针类型绑定

```

 package main

 import (
	 "fmt"
 )

 type A struct {
	 Num int
 }


type Person struct {
	Name string
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

//为了提高效率，通常我们方法和指针类型绑定

func (c *Circle) area2 float{
	c.radius = 10
	return 3.14 * c.radius * c.radius
}
```

# 自定义类型 都可以有方法


Golang 中的方法是作用在指定的数据类型上的(即:和指定的数据类型绑定)，因此自定义类型， 都可以有方法，而不仅仅是 struct。