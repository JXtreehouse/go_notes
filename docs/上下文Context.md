<!--
 * @Author: your name
 * @Date: 2021-04-02 18:24:06
 * @LastEditTime: 2021-04-02 18:25:34
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /go_notes/docs/上下文Context.md
-->
Go 语言中的 [context.Context](https://github.com/golang/go/blob/71bbffbc48d03b447c73da1f54ac57350fc9b36a/src/context/context.go#L62-L154) 的主要作用还是在多个 Goroutine 组成的树中同步取消信号以减少对资源的消耗和占用，虽然它也有传值的功能，但是这个功能我们还是很少用到。