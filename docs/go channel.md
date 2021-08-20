<!--
 * @Author: your name
 * @Date: 2021-07-02 16:16:08
 * @LastEditTime: 2021-07-02 16:21:45
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /go_notes/docs/go channel.md
-->
# channel 通道
通道（channel）是Go语言提供的一种在goroutine之间进行数据传输的通信机制。当然，通过channel传递的数据只能是一些指定的类型，这些类型称为通道的元素类型。此外，要使通道正常运行还需要保证通道有数据接收方。

![](../assets/channel%20example.png)

  通道的声明非常简单，只需要使用chan关键字即可，关闭通道则使用close函数。

  