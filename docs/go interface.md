<!--
 * @Author: your name
 * @Date: 2021-06-24 20:22:49
 * @LastEditTime: 2021-06-24 20:35:01
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * ＠ Reference:  https://github.com/skywind3000/awesome-cheatsheets/blob/master/languages/golang.go
 *   https://github.com/xxjwxc/uber_go_guide_cn#%E6%8C%87%E5%90%91-interface-%E7%9A%84%E6%8C%87%E9%92%88
 * @FilePath: /go_notes/docs/go interface.md
-->

<b>Golang接口</b>


go的接口为鸭子类型，即只要你实现了接口中的方法就实现了该接口

```
type Reader interface {
    Reading()   // 仅需要实现Reading方法就实现了该接口
}

type As struct {}
func (a As) Reading() {}      // 实现了Reader接口

type Bs struct {}
func (b Bs) Reading() {}      // 也实现了Reader接口
func (b Bs) Closing() {}

```

# 指向interface的指针
我们几乎不需要指向接口类型的指针
应该将接口作为值进行传递，在这样的传递过程中，实质上传递的底层数据仍然可以是指针。
接口实质上在底层用两个字段表示：

- 一个指向某些特定类型信息的指针。可以将其视为"type"。
- 数据指针。如果存储的数据是指针，则直接存储。如果存储的数据是一个值，则存储指向该值的指针。
如果希望接口方法修改基础数据，则必须使用指针传递(将对象指针赋值给接口变量)。

```golang
type F interface {
    f()
}

type S1 struct{}
func (s S1) f() {}
type S2 struct{}

func (s *S2) f() {}

// f1.f()无法修改底层数据
// f2.f() 可以修改底层数据,给接口变量f2赋值时使用的是对象指针
var f1 F = S1{}
var f2 F = &S2{}


```




