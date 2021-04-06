<!--
 * @Author: your name
 * @Date: 2021-04-06 15:39:06
 * @LastEditTime: 2021-04-06 20:25:48
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /go_notes/docs/Golang panic用法.md
-->
Go语言追求简洁优雅，所以，Go语言不支持传统的 try…catch…finally 这种异常，因为Go语言的设计者们认为，将异常与控制结构混在一起会很容易使得代码变得混乱。因为开发者很容易滥用异常，甚至一个小小的错误都抛出一个异常。在Go语言中，使用多值返回来返回错误。不要用异常代替错误，更不要用来控制流程。在极个别的情况下，也就是说，遇到真正的异常的情况下（比如除数为 0了）。才使用Go中引入的Exception处理：defer, panic, recover。

这几个异常的使用场景可以这么简单描述：Go中可以抛出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理。

```
package main

import "fmt"

func main(){
    fmt.Println("c")
     defer func(){ // 必须要先声明defer，否则不能捕获到panic异常
        fmt.Println("d")
        if err:=recover();err!=nil{
            fmt.Println(err) // 这里的err其实就是panic传入的内容，55
        }
        fmt.Println("e")
    }()

    f() //开始调用f
    fmt.Println("f") //这里开始下面代码不会再执行
}

func f(){
    fmt.Println("a")
    panic("异常信息")
    fmt.Println("b") //这里开始下面代码不会再执行
    fmt.Println("f")
}


输出结果：

c
a
d
异常信息
e
```

# panic：

- 内建函数
- 假如函数F中书写了panic语句，会终止其后要执行的代码，在panic所在函数F内如果存在要执行的defer函数列表，按照defer的逆序执行
```
比如panic函数内有：
    defer   函数1
    defer   函数2
    defer   函数3
那么执行顺序就是：
    函数3
    函数2
    函数1
```

- 返回函数F的调用者G，在G中，调用函数F语句之后的代码不会执行，假如函数G中存在要执行的defer函数列表，按照defer的逆序执行，这里的defer 有点类似 try-catch-finally 中的 finally
- 到goroutine整个退出，并报告错误
# recover：

- 内建函数
- 用来控制一个goroutine的panicking行为，捕获panic，从而影响应用的行为
- 一般的调用建议(如上面的例子)
a). 在defer函数中，通过recever来终止一个gojroutine的panicking过程，从而恢复正常代码的执行
b). 可以获取通过panic传递的error
简单来讲：go中可以抛出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理。



注意：利用recover处理panic指令，defer必须在panic之前声明，否则当panic时，recover无法捕获到panic．

# 应用示例： 

比如[go编译器的源码](https://github.com/golang/go/blob/master/src/cmd/compile/internal/gc/main.go)
> 注释: 编译器是构建正在运行的二进制文件的编译器工具链的名称。已知的工具链为：
> gc      Also known as cmd/compile.
gccgo   The gccgo front end, part of the GCC compiler suite.

![](../assets/cmd_compile_internal_gc_main.png)