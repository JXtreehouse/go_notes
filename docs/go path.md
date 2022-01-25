<!--
 * @Author: your name
 * @Date: 2021-04-07 11:46:40
 * @LastEditTime: 2021-11-10 15:48:43
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /go_notes/docs/go path.md
-->
# 使用go path问题

- 代码开发必须在go path src目录下，不然，就有问题。
- 依赖手动管理
- 依赖包没有版本可言


假设指定gopath是文件系统的普通目录名，当然我们可以随便设置一个目录名，然后将其路径存入GOPATH。

GOPATH可以是多个目录：


在window系统设置环境变量；在Linux/MacOS系统只要输入终端命令exportgopath=项目所在地址，但是必须保证gopath 这个代码目录下面有三个目录 pkg、bin、src。
新建项目的源码放在src目录下面，以我们现在要写个unzip的go包为例



