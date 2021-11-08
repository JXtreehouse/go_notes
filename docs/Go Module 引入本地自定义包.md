<!--
 * @Author: your name
 * @Date: 2021-11-01 19:32:47
 * @LastEditTime: 2021-11-01 19:37:49
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /go_notes/docs/Go Module 引入本地自定义包.md
-->
# 1. 启用 Go Module

首先在默认情况下，$GOPATH 默认情况下是不支持 go mudules 的，当你执行 go mod init 的时候会遇到如下错误：

```
go: modules disabled inside GOPATH/src by GO111MODULE=auto; see ‘go help modules’
```
我们需要在执行 go mod 命令之前，导出 GO111MODULE 环境变量，你可以直接临时一次性导出， 为了后面方便，建议直接在 ~/.bashrc 文件中导出， 在文件末尾加入：

```
export GO111MODULE=on
```



从这也表明了 go 将来是要利用 modules 机制去消灭 $GOPATH 的。

# 创建 Go Module

我们现在 $GOPATH 下面先创建一个项目，并初始化 module







