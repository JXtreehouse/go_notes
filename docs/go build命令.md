<!--
 * @Author: your name
 * @Date: 2021-04-06 14:41:20
 * @LastEditTime: 2021-04-06 14:47:20
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /go_notes/docs/go build命令.md
-->
Go语言的编译速度非常快。Go 1.9 版本后默认利用Go语言的并发特性进行函数粒度的并发编译。

Go语言的程序编写基本以源码方式，无论是自己的代码还是第三方代码，并且以 GOPATH 作为工作目录和一套完整的工程目录规则。因此Go语言中日常编译时无须像 C++ 一样配置各种包含路径、链接库地址等。

Go语言中使用 go build 命令主要用于编译代码。在包的编译过程中，若有必要，会同时编译与之相关联的包。

go build 有很多种编译方法，如无参数编译、文件列表编译、指定包编译等，使用这些方法都可以输出可执行文件。
http://c.biancheng.net/view/120.html