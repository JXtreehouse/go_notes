<!--
 * @Author: your name
 * @Date: 2021-08-20 11:13:10
 * @LastEditTime: 2021-08-20 14:12:50
 * @LastEditors: Please set LastEditors
 * @Description: build an efficient kv storage engine for Go-based applications,
 * @FilePath: /go_notes/手写系列/手写数据库/实现一个key-value存储引擎.md
-->


我们想设计一个存储，首先考虑的是下面几个问题

- 数据存在哪
- 怎么存（写）
- 怎么取（读）

我们知道，计算机当中有内存和磁盘，内存是易失性的，掉电之后存储的数据全部丢失，所以，如果想要系统崩溃再重启之后依然正常使用，就不得不将数据存储在非易失性介质当中，最常见的便是磁盘。

我们需要设计内存中应该怎么存放，磁盘中应该怎么存放。

主要将数据存储的模型分为了两类：

- B+ 树
- LSM 树


