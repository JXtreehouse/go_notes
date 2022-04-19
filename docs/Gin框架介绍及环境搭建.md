#! https://zhuanlan.zhihu.com/p/478076646
<!--
 * @Author: your name
 * @Date: 2022-03-09 11:31:13
 * @LastEditTime: 2022-03-09 11:58:08
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_notes/docs/Gin框架介绍及环境搭建.md
-->
# Gin简介

Gin是一个golang的微框架，封装比较优雅，API友好, 源代码比较明确。
具有快速灵活，容错方便等特点。其实对于golang而言，web框架的依赖远比Python，Java之类的要小。自身的net/http足够简单，性能也非常不错。框架更像是一个常用函数或者工具的集合。借助框架开发，不仅可以省去很多常用的封装带来的时间，也有助于团队的编码风格和形成规范。

- Gin源码库地址: https://github.com/gin-gonic/gin

- 官网: https://gin-gonic.com/

# Gin特点和特性

- 速度： Gin之所以被很多企业和团队使用，第一个原因是因为速度快，性能表现出众
- 中间件: 和iris类似, gin在处理请求时,支持中间件操作, 方便编码处理
- 路由: gin中可以非常简单的实现路由解析功能，并包含路由组解析功能
- 内置渲染: Gin支持JSON、XML和HTML等多种数据格式的渲染， 并提供了方便的操作API

# 学习文档

官方文档很好用: https://gin-gonic.com/docs/


# Gin开发环境搭建

## 环境要求
gin框架需要go语言版本在1.3及以上。可以通过go verison查看自己的go语言版本是否符合要求
## 安装gin框架

[Quickstart](https://gin-gonic.com/docs/quickstart/)