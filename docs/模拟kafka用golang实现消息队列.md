<!--
 * @Author: your name
 * @Date: 2022-03-04 17:04:19
 * @LastEditTime: 2022-03-04 17:41:40
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_notes/docs/模拟kafka用golang实现消息队列.md
-->
我们要用golang
实现一个类似Kafka的事件总线，它将包括以下功能:

- 集中式的日志收集
- 为机器学习或分析收集结构化或非结构化事件（例如 Hadoop、ClickHouse 等）
- 更改日志缓冲区以使一个或多个存储的内容彼此同步（例如，在数据库更新时更新全文索引）
  
  
消息应该由换行符分隔，尽管在事件格式方面没有其他限制：它可以是 JSON、纯文本、Protobuf（带有 \ 和换行符转义）等。

应该包括以下特性：

- 分布式的，默认使用异步复制
- 显式确认已读取的数据
- 易于配置开箱即用，不会意外丢失数据