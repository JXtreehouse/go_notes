<!--
 * @Author: your name
 * @Date: 2021-04-02 11:37:37
 * @LastEditTime: 2021-04-07 17:51:10
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /go_notes/docs/go vendor 转换go mod.md
-->

# govendor

- 决了包依赖，一个配置文件就管理
- 依赖包全都下载到项目vendor下，每个项目都把有一份。拉取项目时,开始怀疑人生

# go mod
go modules 是 golang 1.11 新加的特性。现在1.12 已经发布了，是时候用起来了。Modules官方定义为：

> 模块是相关Go包的集合。modules是源代码交换和版本控制的单元。 go命令直接支持使用modules，包括记录和解析对其他模块的依赖性。modules替换旧的基于GOPATH的方法来指定在给定构建中使用哪些源文件。

# 转换
1. go mod init [git地址] (生成go.mod)
2. Go Modules 勾选 Enable Go Modules integration
3. 删除vendor包
4. 如果有灰色，把require删去 再次 go mod tidy


