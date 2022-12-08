<!--
 * @Author: your name
 * @Date: 2021-11-08 16:11:13
 * @LastEditTime: 2021-11-08 16:11:13
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_notes/docs/采用Leaf算法结合Go实现双Buffer桶的ID生成器.md
-->
Leaf算法是一种分布式唯一ID生成算法。它通过节点与毫秒级时间戳相结合，来生成唯一的ID。使用双Buffer桶可以提高ID生成器的性能，因为它可以在生成ID的同时进行写入和读取操作。双Buffer桶的实现方式可能因语言而异，但是通常来说，需要使用多线程技术来支持双Buffer桶的工作。例如，在Go中，可以使用channel和goroutine来实现双Buffer桶。


